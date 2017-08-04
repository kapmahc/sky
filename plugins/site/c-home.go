package site

import (
	"net/http"

	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
)

func (p *Plugin) getLocales(c *axe.Context) (interface{}, error) {
	lang := c.Params["lang"]
	return p.I18n.All(lang)
}

func (p *Plugin) getSiteInfo(c *axe.Context) (interface{}, error) {
	// -----------
	langs, err := p.I18n.Store.Languages()
	if err != nil {
		return nil, err
	}
	lng := c.Payload[i18n.LOCALE].(string)
	data := axe.H{"locale": lng, "languages": langs}
	// -----------
	for _, k := range []string{"title", "subTitle", "keywords", "description", "copyright"} {
		data[k], _ = p.I18n.Store.Get(lng, "site."+k)
	}
	// -----------
	author := axe.H{}
	for _, k := range []string{"name", "email"} {
		author[k], _ = p.I18n.Store.Get(lng, "site.author."+k)
	}
	data["author"] = author
	// -----------
	var links []Link
	if err := p.Db.Order("loc DESC, sort_order DESC").Find(&links).Error; err != nil {
		return nil, err
	}
	data["links"] = links
	// -----------
	var cards []Card
	if err := p.Db.Order("loc DESC, sort_order DESC").Find(&cards).Error; err != nil {
		return nil, err
	}
	data["cards"] = cards
	// -----------
	donates := axe.H{}
	paypal := make(map[string]interface{})
	if err := p.Settings.Get("site.paypal", &paypal); err == nil {
		donates["paypal"] = paypal["donate"]
	}
	data["donates"] = donates
	// -----------
	var friendLinks []FriendLink
	if err := p.Db.Order("sort_order DESC").Find(&friendLinks).Error; err == nil {
		data["friendLinks"] = friendLinks
	}
	// -----------
	c.JSON(http.StatusOK, data)
	return nil
}

func (p *Plugin) getDonates(c *axe.Context) (interface{}, error) {
	data := axe.H{}
	var paypal map[string]interface{}
	if err := p.Settings.Get("site.paypal", &paypal); err == nil {
		data["paypal"] = paypal["donate"]
	}
	c.JSON(http.StatusOK, data)
	return nil
}