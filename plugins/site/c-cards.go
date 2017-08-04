package site

import (
	"net/http"

	"github.com/kapmahc/axe"
)

func (p *Plugin) indexCards(c *axe.Context) (interface{}, error) {
	var items []Card
	if err := p.Db.Order("loc ASC, sort_order ASC").Find(&items).Error; err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, items)
	return nil
}

type fmCard struct {
	Loc       string `json:"loc" binding:"required,max=32"`
	Title     string `json:"title" binding:"required,max=255"`
	Summary   string `json:"summary" binding:"required"`
	Type      string `json:"type" binding:"required"`
	Href      string `json:"href" binding:"required,max=255"`
	Logo      string `json:"logo" binding:"required,max=255"`
	Action    string `json:"action" binding:"required,max=32"`
	SortOrder int    `json:"sortOrder"`
}

func (p *Plugin) createCard(c *axe.Context) (interface{}, error) {
	var fm fmCard
	if err := c.Bind(&fm); err != nil {
		return nil, err
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
		return nil, err
	}
	c.JSON(http.StatusOK, item)
	return nil
}

func (p *Plugin) showCard(c *axe.Context) (interface{}, error) {

	var item Card
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, item)
	return nil
}

func (p *Plugin) updateCard(c *axe.Context) (interface{}, error) {
	var fm fmCard
	if err := c.Bind(&fm); err != nil {
		return nil, err
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
		return nil, err
	}
	c.JSON(http.StatusOK, axe.H{})
	return nil
}

func (p *Plugin) destroyCard(c *axe.Context) (interface{}, error) {
	if err := p.Db.
		Where("id = ?", c.Params["id"]).
		Delete(Card{}).Error; err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, axe.H{})
	return nil
}