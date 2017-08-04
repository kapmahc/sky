package mail

import (
	"net/http"

	"github.com/kapmahc/axe"
)

func (p *Plugin) indexDomains(c *axe.Context) {

	var items []Domain
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, items)
}

type fmDomain struct {
	Name string `json:"name" binding:"required,max=255"`
}

func (p *Plugin) createDomain(c *axe.Context) {

	var fm fmDomain
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	item := Domain{
		Name: fm.Name,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (p *Plugin) showDomain(c *axe.Context) {
	var item Domain
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (p *Plugin) updateDomain(c *axe.Context) {
	var fm fmDomain
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	var item Domain
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	if err := p.Db.Model(&item).
		Updates(map[string]interface{}{
			"name": fm.Name,
		}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (p *Plugin) destroyDomain(c *axe.Context) {
	if err := p.Db.
		Where("id = ?", c.Params["id"]).
		Delete(Domain{}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, axe.H{})
}
