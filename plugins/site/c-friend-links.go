package site

import (
	"net/http"

	"github.com/kapmahc/axe"
)

func (p *Plugin) indexFriendLinks(c *axe.Context) (interface{}, error) {
	var items []FriendLink
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, items)
	return nil

}

type fmFriendLink struct {
	Title     string `json:"title" binding:"required,max=255"`
	Home      string `json:"home" binding:"required,max=255"`
	Logo      string `json:"logo" binding:"required,max=255"`
	SortOrder int    `json:"sortOrder"`
}

func (p *Plugin) createFriendLink(c *axe.Context) (interface{}, error) {
	var fm fmFriendLink
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	item := FriendLink{
		Title:     fm.Title,
		Logo:      fm.Logo,
		Home:      fm.Home,
		SortOrder: fm.SortOrder,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, item)
	return nil
}

func (p *Plugin) showFriendLink(c *axe.Context) (interface{}, error) {
	var item FriendLink
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, item)
	return nil
}

func (p *Plugin) updateFriendLink(c *axe.Context) (interface{}, error) {
	var fm fmFriendLink
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	if err := p.Db.Model(&FriendLink{}).
		Where("id = ?", c.Params["id"]).
		Updates(map[string]interface{}{
			"home":       fm.Home,
			"title":      fm.Title,
			"logo":       fm.Logo,
			"sort_order": fm.SortOrder,
		}).Error; err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, axe.H{})
	return nil
}

func (p *Plugin) destroyFriendLink(c *axe.Context) (interface{}, error) {
	if err := p.Db.
		Where("id = ?", c.Params["id"]).
		Delete(FriendLink{}).Error; err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, axe.H{})
	return nil
}
