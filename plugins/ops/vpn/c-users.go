package vpn

import (
	"time"

	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
)

func (p *Plugin) indexUsers(c *axe.Context) (interface{}, error) {

	var items []User
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

type fmUserNew struct {
	FullName             string `json:"fullName" binding:"required,max=255"`
	Email                string `json:"email" binding:"required,email"`
	Password             string `json:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"eqfield=Password"`
	Details              string `json:"details"`
	Enable               bool   `json:"enable"`
	StartUp              string `json:"startUp"`
	ShutDown             string `json:"shutDown"`
}

func (p *Plugin) createUser(c *axe.Context) (interface{}, error) {

	var fm fmUserNew
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	startUp, err := time.Parse(time.RFC3339, fm.StartUp)
	if err != nil {
		return nil, err
	}
	shutDown, err := time.Parse(time.RFC3339, fm.ShutDown)
	if err != nil {
		return nil, err
	}
	user := User{
		FullName: fm.FullName,
		Email:    fm.Email,
		Details:  fm.Details,
		Enable:   fm.Enable,
		StartUp:  startUp,
		ShutDown: shutDown,
	}
	if err := user.SetPassword(fm.Password); err != nil {
		return nil, err
	}
	if err := p.Db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (p *Plugin) showUser(c *axe.Context) (interface{}, error) {
	var item User
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

type fmUserEdit struct {
	FullName string `json:"fullName" binding:"required,max=255"`
	Details  string `json:"details"`
	Enable   bool   `json:"enable"`
	StartUp  string `json:"startUp"`
	ShutDown string `json:"shutDown"`
}

func (p *Plugin) updateUser(c *axe.Context) (interface{}, error) {
	var fm fmUserEdit
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	var item User
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		return nil, err
	}

	startUp, err := time.Parse(time.RFC3339, fm.StartUp)
	if err != nil {
		return nil, err
	}
	shutDown, err := time.Parse(time.RFC3339, fm.ShutDown)
	if err != nil {
		return nil, err
	}
	if err := p.Db.Model(&item).
		Updates(map[string]interface{}{
			"full_name": fm.FullName,
			"enable":    fm.Enable,
			"start_up":  startUp,
			"shut_down": shutDown,
			"details":   fm.Details,
		}).Error; err != nil {
		return nil, err
	}

	return item, nil
}

type fmUserResetPassword struct {
	Password             string `json:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Plugin) postResetUserPassword(c *axe.Context) (interface{}, error) {

	var item User
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		return nil, err
	}
	var fm fmUserResetPassword
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	if err := item.SetPassword(fm.Password); err != nil {
		return nil, err
	}
	if err := p.Db.Model(&item).
		Updates(map[string]interface{}{
			"password": item.Password,
		}).Error; err != nil {
		return nil, err
	}

	return axe.H{}, nil
}

type fmUserChangePassword struct {
	Email                string `json:"email" binding:"required,email"`
	CurrentPassword      string `json:"currentPassword" binding:"required"`
	NewPassword          string `json:"newPassword" binding:"min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"eqfield=NewPassword"`
}

func (p *Plugin) postChangeUserPassword(c *axe.Context) (interface{}, error) {
	lng := c.Payload[i18n.LOCALE].(string)
	var fm fmUserChangePassword
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	var user User
	if err := p.Db.Where("email = ?", fm.Email).First(&user).Error; err != nil {
		return nil, err
	}
	if !user.ChkPassword(fm.CurrentPassword) {
		return nil, p.I18n.E(lng, "ops.vpn.errors.user.email-password-not-match")
	}
	if err := user.SetPassword(fm.NewPassword); err != nil {
		return nil, err
	}

	if err := p.Db.Model(user).
		Updates(map[string]interface{}{
			"password": user.Password,
		}).Error; err != nil {
		return nil, err
	}

	return axe.H{}, nil
}

func (p *Plugin) destroyUser(c *axe.Context) (interface{}, error) {
	if err := p.Db.
		Where("id = ?", c.Params["id"]).
		Delete(User{}).Error; err != nil {
		return nil, err
	}
	return axe.H{}, nil
}
