package site

import (
	"net/http"

	"github.com/kapmahc/axe"
)

func (p *Plugin) indexLeaveWords(c *axe.Context) (interface{}, error) {
	var items []LeaveWord
	if err := p.Db.Order("created_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, items)
	return nil
}

type fmLeaveWord struct {
	Body string `json:"body" binding:"required,max=2048"`
	Type string `json:"type" binding:"required,max=16"`
}

func (p *Plugin) createLeaveWord(c *axe.Context) (interface{}, error) {
	var fm fmLeaveWord
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	item := LeaveWord{
		Body: fm.Body,
		Type: fm.Type,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, axe.H{})
	return nil
}

func (p *Plugin) destroyLeaveWord(c *axe.Context) (interface{}, error) {
	if err := p.Db.
		Where("id = ?", c.Params["id"]).
		Delete(LeaveWord{}).Error; err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, axe.H{})
	return nil
}
