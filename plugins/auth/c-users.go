package auth

import (
	"time"

	"github.com/SermoDigital/jose/jws"
	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
)

type fmSignUp struct {
	Name                 string `json:"name" binding:"required,max=255"`
	Email                string `json:"email" binding:"required,email"`
	Password             string `json:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Plugin) postUsersSignUp(c *axe.Context) (interface{}, error) {
	l := c.Payload[i18n.LOCALE].(string)
	var fm fmSignUp
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	var count int
	if err := p.Db.
		Model(&User{}).
		Where("email = ?", fm.Email).
		Count(&count).Error; err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, p.I18n.E(l, "auth.errors.user.email-already-exists")
	}

	user, err := p.Dao.AddEmailUser(fm.Name, fm.Email, fm.Password)
	if err != nil {
		return nil, err
	}

	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.user.sign-up"))
	p.sendEmail(l, user, actConfirm)

	return axe.H{}, nil
}

type fmSignIn struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RememberMe bool   `json:"rememberMe"`
}

func (p *Plugin) postUsersSignIn(c *axe.Context) (interface{}, error) {
	l := c.Payload[i18n.LOCALE].(string)
	var fm fmSignIn
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	user, err := p.Dao.SignIn(l, fm.Email, fm.Password, c.ClientIP())
	if err != nil {
		return nil, err
	}

	cm := jws.Claims{}
	cm.Set(UID, user.UID)
	cm.Set("admin", p.Dao.Is(user.ID, RoleAdmin))
	tkn, err := p.Jwt.Sum(cm, time.Hour*24*7)
	if err != nil {
		return nil, err
	}

	return axe.H{
		"token": string(tkn),
		"name":  user.Name,
	}, nil
}

type fmEmail struct {
	Email string `json:"email" binding:"required,email"`
}

func (p *Plugin) getUsersConfirm(c *axe.Context) (interface{}, error) {
	l := c.Payload[i18n.LOCALE].(string)
	user, err := p.parseToken(l, c.Params["token"], actConfirm)
	if err != nil {
		return nil, err
	}
	if user.IsConfirm() {
		return nil, p.I18n.E(l, "auth.errors.user.already-confirm")
	}
	p.Db.Model(user).Update("confirmed_at", time.Now())
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.user.confirm"))

	return axe.H{}, nil
}

func (p *Plugin) postUsersConfirm(c *axe.Context) (interface{}, error) {
	l := c.Payload[i18n.LOCALE].(string)
	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return nil, err
	}

	if user.IsConfirm() {
		return nil, p.I18n.E(l, "auth.errors.user.already-confirm")
	}

	p.sendEmail(l, user, actConfirm)
	return axe.H{}, nil
}

func (p *Plugin) getUsersUnlock(c *axe.Context) (interface{}, error) {
	l := c.Payload[i18n.LOCALE].(string)
	user, err := p.parseToken(l, c.Params["token"], actUnlock)
	if err != nil {
		return nil, err
	}
	if !user.IsLock() {
		return nil, p.I18n.E(l, "auth.errors.user.not-lock")
	}

	if err := p.Db.Model(user).Update(map[string]interface{}{"locked_at": nil}).Error; err != nil {
		return nil, err
	}
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.user.unlock"))

	return axe.H{}, nil
}

func (p *Plugin) postUsersUnlock(c *axe.Context) (interface{}, error) {
	l := c.Payload[i18n.LOCALE].(string)

	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	if !user.IsLock() {
		return nil, p.I18n.E(l, "auth.errors.user.not-lock")
	}
	p.sendEmail(l, user, actUnlock)
	return axe.H{}, nil
}

func (p *Plugin) postUsersForgotPassword(c *axe.Context) (interface{}, error) {
	l := c.Payload[i18n.LOCALE].(string)
	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	var user *User
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	p.sendEmail(l, user, actResetPassword)

	return axe.H{}, nil
}

type fmResetPassword struct {
	Token                string `json:"token" binding:"required"`
	Password             string `json:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Plugin) postUsersResetPassword(c *axe.Context) (interface{}, error) {
	l := c.Payload[i18n.LOCALE].(string)

	var fm fmResetPassword
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user, err := p.parseToken(l, fm.Token, actResetPassword)
	if err != nil {
		return nil, err
	}
	p.Db.Model(user).Update("password", p.Hmac.Sum([]byte(fm.Password)))
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.user.reset-password"))

	return axe.H{}, nil
}

func (p *Plugin) deleteUsersSignOut(c *axe.Context) (interface{}, error) {
	l := c.Payload[i18n.LOCALE].(string)
	user := c.Payload[CurrentUser].(*User)
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.user.sign-out"))

	return axe.H{}, nil
}

func (p *Plugin) getUsersInfo(c *axe.Context) (interface{}, error) {
	user := c.Payload[CurrentUser].(*User)
	return axe.H{"name": user.Name, "email": user.Email}, nil
}

type fmInfo struct {
	Name string `json:"name" binding:"required,max=255"`
	// Home string `json:"home" binding:"max=255"`
	// Logo string `json:"logo" binding:"max=255"`
}

func (p *Plugin) postUsersInfo(c *axe.Context) (interface{}, error) {
	user := c.Payload[CurrentUser].(*User)
	var fm fmInfo
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	if err := p.Db.Model(user).Updates(map[string]interface{}{
		// "home": fm.Home,
		// "logo": fm.Logo,
		"name": fm.Name,
	}).Error; err != nil {
		return nil, err
	}
	return axe.H{}, nil
}

type fmChangePassword struct {
	CurrentPassword      string `json:"currentPassword" binding:"required"`
	NewPassword          string `json:"newPassword" binding:"min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"eqfield=NewPassword"`
}

func (p *Plugin) postUsersChangePassword(c *axe.Context) (interface{}, error) {
	l := c.Payload[i18n.LOCALE].(string)

	user := c.Payload[CurrentUser].(*User)
	var fm fmChangePassword
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	if !p.Hmac.Chk([]byte(fm.CurrentPassword), user.Password) {
		return nil, p.I18n.E(l, "auth.errors.bad-password")
	}
	if err := p.Db.Model(user).
		Update("password", p.Hmac.Sum([]byte(fm.NewPassword))).Error; err != nil {
		return nil, err
	}

	return axe.H{}, nil
}

func (p *Plugin) getUsersLogs(c *axe.Context) (interface{}, error) {
	user := c.Payload[CurrentUser].(*User)
	var items []Log
	if err := p.Db.
		Select([]string{"id", "ip", "message", "created_at"}).
		Where("user_id = ?", user.ID).
		Order("id DESC").Limit(120).
		Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}
