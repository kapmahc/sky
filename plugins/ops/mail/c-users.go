package mail

import (
	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
)

func (p *Plugin) indexUsers(c *axe.Context) (interface{}, error) {
	var items []User
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}
	var domains []Domain
	if err := p.Db.Select([]string{"id", "name"}).Find(&domains).Error; err != nil {
		return nil, err
	}
	for i := range items {
		u := &items[i]
		for _, d := range domains {
			if d.ID == u.DomainID {
				u.Domain = d
				break
			}
		}
	}

	return items, nil
}

type fmUserNew struct {
	FullName             string `json:"fullName" binding:"required,max=255"`
	Email                string `json:"email" binding:"required,email"`
	Password             string `json:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"eqfield=Password"`
	Enable               bool   `json:"enable"`
	DomainID             uint   `json:"domainId"`
}

func (p *Plugin) createUser(c *axe.Context) (interface{}, error) {

	var fm fmUserNew
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user := User{
		FullName: fm.FullName,
		Email:    fm.Email,
		Enable:   fm.Enable,
		DomainID: fm.DomainID,
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
	Enable   bool   `json:"enable"`
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

	if err := p.Db.Model(&item).
		Updates(map[string]interface{}{
			"enable":    fm.Enable,
			"full_name": fm.FullName,
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

	var fm fmUserResetPassword
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	var item User
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
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
		return nil, p.I18n.E(lng, "ops.mail.errors.user.email-password-not-match")
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
	lng := c.Payload[i18n.LOCALE].(string)
	var user User
	if err := p.Db.
		Where("id = ?", c.Params["id"]).First(&user).Error; err != nil {
		return nil, err
	}
	var count int
	if err := p.Db.Model(&Alias{}).Where("destination = ?", user.Email).Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, p.I18n.E(lng, "errors.in-use")
	}
	if err := p.Db.Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
