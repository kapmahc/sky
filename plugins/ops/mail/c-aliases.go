package mail

import (
	"net/http"

	"github.com/kapmahc/axe"
)

func (p *Plugin) indexAliases(c *axe.Context) {

	var items []Alias
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	var domains []Domain
	if err := p.Db.Select([]string{"id", "name"}).Find(&domains).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
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

	c.JSON(http.StatusOK, items)
}

type fmAlias struct {
	Source      string `json:"source" validate:"required,max=255"`
	Destination string `json:"destination" validate:"required,max=255"`
}

func (p *Plugin) createAlias(c *axe.Context) {
	var fm fmAlias
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	var user User
	if err := p.Db.Where("email = ?", fm.Destination).First(&user).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	item := Alias{
		Source:      fm.Source,
		Destination: fm.Destination,
		DomainID:    user.DomainID,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (p *Plugin) showAlias(c *axe.Context) {
	var item Alias
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (p *Plugin) updateAlias(c *axe.Context) {
	var fm fmAlias
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	var item Alias
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	var user User
	if err := p.Db.Where("email = ?", fm.Destination).First(&user).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	if err := p.Db.Model(&item).
		Updates(map[string]interface{}{
			"domain_id":   user.DomainID,
			"source":      fm.Source,
			"destination": fm.Destination,
		}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (p *Plugin) destroyAlias(c *axe.Context) {
	if err := p.Db.
		Where("id = ?", c.Params["id"]).
		Delete(Alias{}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, axe.H{})
}
