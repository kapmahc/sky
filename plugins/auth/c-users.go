package auth

import (
	"net/http"
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

func (p *Plugin) postUsersSignUp(c *axe.Context) {
	l := c.Payload[i18n.LOCALE].(string)
	var fm fmSignUp
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	var count int
	if err := p.Db.
		Model(&User{}).
		Where("email = ?", fm.Email).
		Count(&count).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	if count > 0 {
		c.Abort(http.StatusInternalServerError, p.I18n.E(l, "auth.errors.user.email-already-exists"))
		return
	}

	user, err := p.Dao.AddEmailUser(fm.Name, fm.Email, fm.Password)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.user.sign-up"))
	p.sendEmail(l, user, actConfirm)

	c.JSON(http.StatusOK, axe.H{})
}

type fmSignIn struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RememberMe bool   `json:"rememberMe"`
}

func (p *Plugin) postUsersSignIn(c *axe.Context) {
	l := c.Payload[i18n.LOCALE].(string)
	var fm fmSignIn
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	user, err := p.Dao.SignIn(l, fm.Email, fm.Password, c.ClientIP())
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	cm := jws.Claims{}
	cm.Set(UID, user.UID)
	cm.Set("admin", p.Dao.Is(user.ID, RoleAdmin))
	tkn, err := p.Jwt.Sum(cm, time.Hour*24*7)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, axe.H{
		"token": string(tkn),
		"name":  user.Name,
	})
}

type fmEmail struct {
	Email string `json:"email" binding:"required,email"`
}

func (p *Plugin) getUsersConfirm(c *axe.Context) {
	l := c.Payload[i18n.LOCALE].(string)
	user, err := p.parseToken(l, c.Params["token"], actConfirm)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if user.IsConfirm() {
		c.Abort(http.StatusInternalServerError, p.I18n.E(l, "auth.errors.user.already-confirm"))
		return
	}
	p.Db.Model(user).Update("confirmed_at", time.Now())
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.user.confirm"))

	c.JSON(http.StatusOK, axe.H{})
}

func (p *Plugin) postUsersConfirm(c *axe.Context) {
	l := c.Payload[i18n.LOCALE].(string)
	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	if user.IsConfirm() {
		c.Abort(http.StatusInternalServerError, p.I18n.E(l, "auth.errors.user.already-confirm"))
		return
	}

	p.sendEmail(l, user, actConfirm)
	c.JSON(http.StatusOK, axe.H{})
}

func (p *Plugin) getUsersUnlock(c *axe.Context) {
	l := c.Payload[i18n.LOCALE].(string)
	user, err := p.parseToken(l, c.Params["token"], actUnlock)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if !user.IsLock() {
		c.Abort(http.StatusInternalServerError, p.I18n.E(l, "auth.errors.user.not-lock"))
		return
	}

	if err := p.Db.Model(user).Update(map[string]interface{}{"locked_at": nil}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.user.unlock"))

	c.JSON(http.StatusOK, axe.H{})
}

func (p *Plugin) postUsersUnlock(c *axe.Context) {
	l := c.Payload[i18n.LOCALE].(string)

	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if !user.IsLock() {
		c.Abort(http.StatusInternalServerError, p.I18n.E(l, "auth.errors.user.not-lock"))
		return
	}
	p.sendEmail(l, user, actUnlock)
	c.JSON(http.StatusOK, axe.H{})
}

func (p *Plugin) postUsersForgotPassword(c *axe.Context) {
	l := c.Payload[i18n.LOCALE].(string)
	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	var user *User
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	p.sendEmail(l, user, actResetPassword)

	c.JSON(http.StatusOK, axe.H{})
}

type fmResetPassword struct {
	Token                string `json:"token" binding:"required"`
	Password             string `json:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Plugin) postUsersResetPassword(c *axe.Context) {
	l := c.Payload[i18n.LOCALE].(string)

	var fm fmResetPassword
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	user, err := p.parseToken(l, fm.Token, actResetPassword)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	p.Db.Model(user).Update("password", p.Hmac.Sum([]byte(fm.Password)))
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.user.reset-password"))

	c.JSON(http.StatusOK, axe.H{})
}

func (p *Plugin) deleteUsersSignOut(c *axe.Context) {
	l := c.Payload[i18n.LOCALE].(string)
	user := c.Payload[CurrentUser].(*User)
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.user.sign-out"))

	c.JSON(http.StatusOK, axe.H{})
}

func (p *Plugin) getUsersInfo(c *axe.Context) {
	user := c.Payload[CurrentUser].(*User)
	c.JSON(http.StatusOK, user)
}

type fmInfo struct {
	Name string `json:"name" binding:"required,max=255"`
	// Home string `json:"home" binding:"max=255"`
	// Logo string `json:"logo" binding:"max=255"`
}

func (p *Plugin) postUsersInfo(c *axe.Context) {
	user := c.Payload[CurrentUser].(*User)
	var fm fmInfo
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	if err := p.Db.Model(user).Updates(map[string]interface{}{
		// "home": fm.Home,
		// "logo": fm.Logo,
		"name": fm.Name,
	}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, axe.H{})
}

type fmChangePassword struct {
	CurrentPassword      string `json:"currentPassword" binding:"required"`
	NewPassword          string `json:"newPassword" binding:"min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"eqfield=NewPassword"`
}

func (p *Plugin) postUsersChangePassword(c *axe.Context) {
	l := c.Payload[i18n.LOCALE].(string)

	user := c.Payload[CurrentUser].(*User)
	var fm fmChangePassword
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if !p.Hmac.Chk([]byte(fm.CurrentPassword), user.Password) {
		c.Abort(http.StatusInternalServerError, p.I18n.E(l, "auth.errors.bad-password"))
		return
	}
	if err := p.Db.Model(user).
		Update("password", p.Hmac.Sum([]byte(fm.NewPassword))).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, axe.H{})
}

func (p *Plugin) getUsersLogs(c *axe.Context) {
	user := c.Payload[CurrentUser].(*User)
	var items []Log
	if err := p.Db.
		Select([]string{"id", "ip", "message", "created_at"}).
		Where("user_id = ?", user.ID).
		Order("id DESC").Limit(120).
		Find(&items).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, items)
}
