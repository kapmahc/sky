package mail

import (
	"github.com/kapmahc/axe"
)

func (p *Plugin) indexDomains(c *axe.Context) (interface{}, error) {

	var items []Domain
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

type fmDomain struct {
	Name string `json:"name" binding:"required,max=255"`
}

func (p *Plugin) createDomain(c *axe.Context) (interface{}, error) {

	var fm fmDomain
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	item := Domain{
		Name: fm.Name,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (p *Plugin) showDomain(c *axe.Context) (interface{}, error) {
	var item Domain
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (p *Plugin) updateDomain(c *axe.Context) (interface{}, error) {
	var fm fmDomain
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	var item Domain
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		return nil, err
	}

	if err := p.Db.Model(&item).
		Updates(map[string]interface{}{
			"name": fm.Name,
		}).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (p *Plugin) destroyDomain(c *axe.Context) (interface{}, error) {
	if err := p.Db.
		Where("id = ?", c.Params["id"]).
		Delete(Domain{}).Error; err != nil {
		return nil, err
	}

	return axe.H{}, nil
}
