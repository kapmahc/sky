package vpn

import (
	"net/http"
	"time"

	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
)

type fmSignIn struct {
	Email    string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"min=6,max=32"`
}

func (p *Plugin) apiAuth(c *axe.Context) {
	var fm fmSignIn
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	lng := c.Payload[i18n.LOCALE].(string)
	var user User
	if err := p.Db.Where("email = ?", fm.Email).First(&user).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	now := time.Now()
	if user.Enable && user.StartUp.Before(now) && user.ShutDown.After(now) {
		c.JSON(http.StatusOK, axe.H{})
		return
	}
	c.Abort(http.StatusInternalServerError, p.I18n.E(lng, "ops.vpn.errors.user.not-available"))
}

type fmStatus struct {
	Email       string  `json:"common_name" validate:"required,email"`
	TrustedIP   string  `json:"trusted_ip" validate:"required"`
	TrustedPort uint    `json:"trusted_port" validate:"required"`
	RemoteIP    string  `json:"ifconfig_pool_remote_ip" validate:"required"`
	RemotePort  uint    `json:"remote_port_1" validate:"required"`
	Received    float64 `json:"bytes_received" validate:"required"`
	Send        float64 `json:"bytes_sent" validate:"required"`
}

func (p *Plugin) apiConnect(c *axe.Context) {
	var fm fmStatus
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	var user User
	if err := p.Db.Where("email = ?", fm.Email).First(&user).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Create(&Log{
		UserID:      user.ID,
		RemoteIP:    fm.RemoteIP,
		RemotePort:  fm.RemotePort,
		TrustedIP:   fm.TrustedIP,
		TrustedPort: fm.TrustedPort,
		Received:    fm.Received,
		Send:        fm.Send,
		StartUp:     time.Now(),
	}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Model(&User{}).
		Where("id = ?", user.ID).
		UpdateColumn("online", true).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, axe.H{})
}

func (p *Plugin) apiDisconnect(c *axe.Context) {
	var fm fmStatus
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	var user User
	if err := p.Db.Where("email = ?", fm.Email).First(&user).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Model(&User{}).
		Where("id = ?", user.ID).
		UpdateColumn("online", false).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	if err := p.Db.
		Model(&Log{}).
		Where(
			"trusted_ip = ? AND trusted_port = ? AND user_id = ? AND shut_down IS NULL",
			fm.TrustedIP,
			fm.TrustedPort,
			user.ID,
		).Update(map[string]interface{}{
		"shut_down": time.Now(),
		"received":  fm.Received,
		"send":      fm.Send,
	}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, axe.H{})
}
