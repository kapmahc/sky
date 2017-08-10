package site

import (
	"net/http"

	"github.com/kapmahc/axe"
)

func (p *Plugin) indexCards(c *axe.Context) {
	var items []Card
	if err := p.Db.Order("loc ASC, sort_order ASC").Find(&items).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, items)
}

type fmCard struct {
	Loc       string `json:"loc" validate:"required,max=32"`
	Title     string `json:"title" validate:"required,max=255"`
	Summary   string `json:"summary" validate:"required"`
	Type      string `json:"type" validate:"required"`
	Href      string `json:"href" validate:"required,max=255"`
	Logo      string `json:"logo" validate:"required,max=255"`
	Action    string `json:"action" validate:"required,max=32"`
	SortOrder int    `json:"sortOrder"`
}

func (p *Plugin) createCard(c *axe.Context) {
	var fm fmCard
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	item := Card{
		Title:     fm.Title,
		Logo:      fm.Logo,
		Href:      fm.Href,
		Summary:   fm.Summary,
		Type:      fm.Type,
		Action:    fm.Action,
		SortOrder: fm.SortOrder,
		Loc:       fm.Loc,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (p *Plugin) showCard(c *axe.Context) {

	var item Card
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (p *Plugin) updateCard(c *axe.Context) {
	var fm fmCard
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Model(&Card{}).
		Where("id = ?", c.Params["id"]).
		Updates(map[string]interface{}{
			"href":       fm.Href,
			"title":      fm.Title,
			"logo":       fm.Logo,
			"sort_order": fm.SortOrder,
			"loc":        fm.Loc,
			"summary":    fm.Summary,
			"type":       fm.Type,
			"action":     fm.Action,
		}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, axe.H{})
}

func (p *Plugin) destroyCard(c *axe.Context) {
	if err := p.Db.
		Where("id = ?", c.Params["id"]).
		Delete(Card{}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, axe.H{})
}
