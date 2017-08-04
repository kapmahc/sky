package site

import (
	"net/http"

	"github.com/kapmahc/axe"
)

func (p *Plugin) indexLinks(c *axe.Context) (interface{}, error) {
	var items []Link
	if err := p.Db.Order("loc ASC, sort_order ASC").Find(&items).Error; err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, items)
	return nil
}

type fmLink struct {
	Label     string `json:"label" binding:"required,max=255"`
	Href      string `json:"href" binding:"required,max=255"`
	Loc       string `json:"loc" binding:"required,max=32"`
	SortOrder int    `json:"sortOrder"`
}

func (p *Plugin) createLink(c *axe.Context) (interface{}, error) {
	var fm fmLink
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	item := Link{
		Label:     fm.Label,
		Href:      fm.Href,
		Loc:       fm.Loc,
		SortOrder: fm.SortOrder,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, item)
	return nil
}

func (p *Plugin) showLink(c *axe.Context) (interface{}, error) {
	var item Link
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, item)
	return nil
}

func (p *Plugin) updateLink(c *axe.Context) (interface{}, error) {
	var fm fmLink
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	if err := p.Db.Model(&Link{}).
		Where("id = ?", c.Params["id"]).
		Updates(map[string]interface{}{
			"loc":        fm.Loc,
			"label":      fm.Label,
			"href":       fm.Href,
			"sort_order": fm.SortOrder,
		}).Error; err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, axe.H{})
	return nil
}

func (p *Plugin) destroyLink(c *axe.Context) (interface{}, error) {
	if err := p.Db.
		Where("id = ?", c.Params["id"]).
		Delete(Link{}).Error; err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, axe.H{})
	return nil
}
