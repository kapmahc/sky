package mail

import (
	"github.com/kapmahc/axe"
)

func (p *Plugin) indexAliases(c *axe.Context) (interface{}, error) {

	var items []Alias
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}

	var domains []Domain
	if err := p.Db.Select([]string{"id", "name"}).Find(&domains).Error; err != nil {
		return nil, err
	}
	for i := range items {
		u := &items[i]
		for _, d := range domains {
			if d.ID == u.DomainID {
				u.Domain = d
				break
			}
		}
	}

	return items, nil
}

type fmAlias struct {
	Source      string `json:"source" binding:"required,max=255"`
	Destination string `json:"destination" binding:"required,max=255"`
}

func (p *Plugin) createAlias(c *axe.Context) (interface{}, error) {
	var fm fmAlias
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	var user User
	if err := p.Db.Where("email = ?", fm.Destination).First(&user).Error; err != nil {
		return nil, err
	}
	item := Alias{
		Source:      fm.Source,
		Destination: fm.Destination,
		DomainID:    user.DomainID,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (p *Plugin) showAlias(c *axe.Context) (interface{}, error) {
	var item Alias
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (p *Plugin) updateAlias(c *axe.Context) (interface{}, error) {
	var fm fmAlias
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	var item Alias
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		return nil, err
	}

	var user User
	if err := p.Db.Where("email = ?", fm.Destination).First(&user).Error; err != nil {
		return nil, err
	}

	if err := p.Db.Model(&item).
		Updates(map[string]interface{}{
			"domain_id":   user.DomainID,
			"source":      fm.Source,
			"destination": fm.Destination,
		}).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (p *Plugin) destroyAlias(c *axe.Context) (interface{}, error) {
	if err := p.Db.
		Where("id = ?", c.Params["id"]).
		Delete(Alias{}).Error; err != nil {
		return nil, err
	}

	return axe.H{}, nil
}
