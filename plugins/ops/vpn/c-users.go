package vpn

import (
	"net/http"
	"time"

	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
)

func (p *Plugin) indexUsers(c *axe.Context) {

	var items []User
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, items)
}

type fmUserNew struct {
	FullName             string `json:"fullName" validate:"required,max=255"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"eqfield=Password"`
	Details              string `json:"details"`
	Enable               bool   `json:"enable"`
	StartUp              string `json:"startUp"`
	ShutDown             string `json:"shutDown"`
}

func (p *Plugin) createUser(c *axe.Context) {

	var fm fmUserNew
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	startUp, err := time.Parse(time.RFC3339, fm.StartUp)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	shutDown, err := time.Parse(time.RFC3339, fm.ShutDown)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
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
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Create(&user).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (p *Plugin) showUser(c *axe.Context) {
	var item User
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, item)
}

type fmUserEdit struct {
	FullName string `json:"fullName" validate:"required,max=255"`
	Details  string `json:"details"`
	Enable   bool   `json:"enable"`
	StartUp  string `json:"startUp"`
	ShutDown string `json:"shutDown"`
}

func (p *Plugin) updateUser(c *axe.Context) {
	var fm fmUserEdit
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	var item User
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	startUp, err := time.Parse(time.RFC3339, fm.StartUp)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	shutDown, err := time.Parse(time.RFC3339, fm.ShutDown)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Model(&item).
		Updates(map[string]interface{}{
			"full_name": fm.FullName,
			"enable":    fm.Enable,
			"start_up":  startUp,
			"shut_down": shutDown,
			"details":   fm.Details,
		}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

type fmUserResetPassword struct {
	Password             string `json:"password" validate:"min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"eqfield=Password"`
}

func (p *Plugin) postResetUserPassword(c *axe.Context) {

	var item User
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	var fm fmUserResetPassword
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	if err := item.SetPassword(fm.Password); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Model(&item).
		Updates(map[string]interface{}{
			"password": item.Password,
		}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, axe.H{})
}

type fmUserChangePassword struct {
	Email                string `json:"email" validate:"required,email"`
	CurrentPassword      string `json:"currentPassword" validate:"required"`
	NewPassword          string `json:"newPassword" validate:"min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"eqfield=NewPassword"`
}

func (p *Plugin) postChangeUserPassword(c *axe.Context) {
	lng := c.Payload[i18n.LOCALE].(string)
	var fm fmUserChangePassword
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	var user User
	if err := p.Db.Where("email = ?", fm.Email).First(&user).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if !user.ChkPassword(fm.CurrentPassword) {
		c.Abort(http.StatusInternalServerError, p.I18n.E(lng, "user.email-password-not-match"))
		return
	}
	if err := user.SetPassword(fm.NewPassword); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	if err := p.Db.Model(user).
		Updates(map[string]interface{}{
			"password": user.Password,
		}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, axe.H{})
}

func (p *Plugin) destroyUser(c *axe.Context) {
	if err := p.Db.
		Where("id = ?", c.Params["id"]).
		Delete(User{}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, axe.H{})
}
