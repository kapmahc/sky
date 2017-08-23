package site

import (
	"net/http"
	"time"

	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
	"github.com/kapmahc/sky/plugins/auth"
)

type fmInstall struct {
	Title                string `json:"title" validate:"required"`
	SubTitle             string `json:"subTitle" validate:"required"`
	Name                 string `json:"name" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"eqfield=Password"`
}

func (p *Plugin) postInstall(c *axe.Context) {
	if err := p._mustDatabaseEmpty(c); err != nil {
		c.Abort(http.StatusForbidden, err)
		return
	}
	var fm fmInstall
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	lng := c.Payload[i18n.LOCALE].(string)
	p.I18n.Set(lng, "site.title", fm.Title)
	p.I18n.Set(lng, "site.subTitle", fm.SubTitle)
	user, err := p.Dao.AddEmailUser(fm.Name, fm.Email, fm.Password)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	for _, r := range []string{auth.RoleAdmin, auth.RoleRoot} {
		role, er := p.Dao.Role(r, auth.DefaultResourceType, auth.DefaultResourceID)
		if er != nil {
			c.Abort(http.StatusInternalServerError, er)
			return
		}
		if err = p.Dao.Allow(role.ID, user.ID, 50, 0, 0); err != nil {
			c.Abort(http.StatusInternalServerError, err)
			return
		}
	}
	if err = p.Db.Model(user).UpdateColumn("confirmed_at", time.Now()).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, axe.H{})
}

func (p *Plugin) _mustDatabaseEmpty(c *axe.Context) error {
	lng := c.Payload[i18n.LOCALE].(string)
	var count int
	if err := p.Db.Model(&auth.User{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return p.I18n.E(lng, "errors.forbidden")
	}
	return nil
}
