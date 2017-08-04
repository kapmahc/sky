package site

import (
	"net/http"

	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
)

func (p *Plugin) indexAdminLocales(c *axe.Context) (interface{}, error) {
	lng := c.Payload[i18n.LOCALE].(string)
	items, err := p.I18n.Store.All(lng)
	if err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, items)
	return nil
}
func (p *Plugin) showAdminLocale(c *axe.Context) (interface{}, error) {
	lng := c.Payload[i18n.LOCALE].(string)
	code := c.Params["code"]
	message, err := p.I18n.Store.Get(lng, code)
	if err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, axe.H{
		"code":    code,
		"message": message,
	})
	return nil
}

func (p *Plugin) destroyAdminLocale(c *axe.Context) (interface{}, error) {
	lng := c.Payload[i18n.LOCALE].(string)
	if err := p.I18n.Store.Del(lng, c.Params["code"]); err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, axe.H{})
	return nil
}

type fmLocale struct {
	Code    string `json:"code" binding:"required,max=255"`
	Message string `json:"message" binding:"required"`
}

func (p *Plugin) postAdminLocales(c *axe.Context) (interface{}, error) {
	var fm fmLocale
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	lng := c.Payload[i18n.LOCALE].(string)
	if err := p.I18n.Set(lng, fm.Code, fm.Message); err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, axe.H{})
	return nil
}
