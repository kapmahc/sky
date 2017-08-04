package vpn

import (
	"time"

	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
)

type fmSignIn struct {
	Email    string `json:"username" binding:"required,email"`
	Password string `json:"password" binding:"min=6,max=32"`
}

func (p *Plugin) apiAuth(c *axe.Context) (interface{}, error) {
	var fm fmSignIn
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	lng := c.Payload[i18n.LOCALE].(string)
	var user User
	if err := p.Db.Where("email = ?", fm.Email).First(&user).Error; err != nil {
		return nil, err
	}
	now := time.Now()
	if user.Enable && user.StartUp.Before(now) && user.ShutDown.After(now) {
		return axe.H{}, nil
	}
	return nil, p.I18n.E(lng, "ops.vpn.errors.user.not-available")
}

type fmStatus struct {
	Email       string  `json:"common_name" binding:"required,email"`
	TrustedIP   string  `json:"trusted_ip" binding:"required"`
	TrustedPort uint    `json:"trusted_port" binding:"required"`
	RemoteIP    string  `json:"ifconfig_pool_remote_ip" binding:"required"`
	RemotePort  uint    `json:"remote_port_1" binding:"required"`
	Received    float64 `json:"bytes_received" binding:"required"`
	Send        float64 `json:"bytes_sent" binding:"required"`
}

func (p *Plugin) apiConnect(c *axe.Context) (interface{}, error) {
	var fm fmStatus
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	var user User
	if err := p.Db.Where("email = ?", fm.Email).First(&user).Error; err != nil {
		return nil, err
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
		return nil, err
	}
	if err := p.Db.Model(&User{}).
		Where("id = ?", user.ID).
		UpdateColumn("online", true).Error; err != nil {
		return nil, err
	}
	return axe.H{}, nil
}

func (p *Plugin) apiDisconnect(c *axe.Context) (interface{}, error) {
	var fm fmStatus
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	var user User
	if err := p.Db.Where("email = ?", fm.Email).First(&user).Error; err != nil {
		return nil, err
	}
	if err := p.Db.Model(&User{}).
		Where("id = ?", user.ID).
		UpdateColumn("online", false).Error; err != nil {
		return nil, err
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
		return nil, err
	}

	return axe.H{}, nil
}