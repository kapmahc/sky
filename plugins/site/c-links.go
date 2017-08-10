package site

import (
	"net/http"

	"github.com/kapmahc/axe"
)

func (p *Plugin) indexLinks(c *axe.Context) {
	var items []Link
	if err := p.Db.Order("loc ASC, sort_order ASC").Find(&items).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, items)
}

type fmLink struct {
	Label     string `json:"label" validate:"required,max=255"`
	Href      string `json:"href" validate:"required,max=255"`
	Loc       string `json:"loc" validate:"required,max=32"`
	SortOrder int    `json:"sortOrder"`
}

func (p *Plugin) createLink(c *axe.Context) {
	var fm fmLink
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	item := Link{
		Label:     fm.Label,
		Href:      fm.Href,
		Loc:       fm.Loc,
		SortOrder: fm.SortOrder,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (p *Plugin) showLink(c *axe.Context) {
	var item Link
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (p *Plugin) updateLink(c *axe.Context) {
	var fm fmLink
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Model(&Link{}).
		Where("id = ?", c.Params["id"]).
		Updates(map[string]interface{}{
			"loc":        fm.Loc,
			"label":      fm.Label,
			"href":       fm.Href,
			"sort_order": fm.SortOrder,
		}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, axe.H{})
}

func (p *Plugin) destroyLink(c *axe.Context) {
	if err := p.Db.
		Where("id = ?", c.Params["id"]).
		Delete(Link{}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, axe.H{})
}
