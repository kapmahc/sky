package site

import (
	"net/http"

	"github.com/kapmahc/axe"
)

func (p *Plugin) indexFriendLinks(c *axe.Context) {
	var items []FriendLink
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, items)

}

type fmFriendLink struct {
	Title     string `json:"title" validate:"required,max=255"`
	Home      string `json:"home" validate:"required,max=255"`
	Logo      string `json:"logo" validate:"required,max=255"`
	SortOrder int    `json:"sortOrder"`
}

func (p *Plugin) createFriendLink(c *axe.Context) {
	var fm fmFriendLink
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	item := FriendLink{
		Title:     fm.Title,
		Logo:      fm.Logo,
		Home:      fm.Home,
		SortOrder: fm.SortOrder,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (p *Plugin) showFriendLink(c *axe.Context) {
	var item FriendLink
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (p *Plugin) updateFriendLink(c *axe.Context) {
	var fm fmFriendLink
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Model(&FriendLink{}).
		Where("id = ?", c.Params["id"]).
		Updates(map[string]interface{}{
			"home":       fm.Home,
			"title":      fm.Title,
			"logo":       fm.Logo,
			"sort_order": fm.SortOrder,
		}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, axe.H{})
}

func (p *Plugin) destroyFriendLink(c *axe.Context) {
	if err := p.Db.
		Where("id = ?", c.Params["id"]).
		Delete(FriendLink{}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, axe.H{})
}
