package site

import (
	"net/http"

	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
)

func (p *Plugin) indexAdminLocale(c *axe.Context) {
	lng := c.Payload[i18n.LOCALE].(string)
	// var items []i_orm.Model
	// if err := p.Db.Select([]string{"code", "message"}).Where("lang = ?", lng).Order("code ASC").Find(&items).Error; err != nil {
	// 	c.Abort(http.StatusInternalServerError, err)
	// 	return
	// }
	items, err := p.I18n.Store.All(lng)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, items)
}

func (p *Plugin) showAdminLocale(c *axe.Context) {
	lng := c.Payload[i18n.LOCALE].(string)
	code := c.Params["code"]
	message, err := p.I18n.Store.Get(lng, code)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, axe.H{
		"code":    code,
		"message": message,
	})
}

func (p *Plugin) destroyAdminLocale(c *axe.Context) {
	lng := c.Payload[i18n.LOCALE].(string)
	if err := p.I18n.Store.Del(lng, c.Params["code"]); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, axe.H{})
}

type fmLocale struct {
	Code    string `json:"code" validate:"required,max=255"`
	Message string `json:"message" validate:"required"`
}

func (p *Plugin) postAdminLocales(c *axe.Context) {
	var fm fmLocale
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	lng := c.Payload[i18n.LOCALE].(string)
	if err := p.I18n.Set(lng, fm.Code, fm.Message); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, axe.H{})
}
