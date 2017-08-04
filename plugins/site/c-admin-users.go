package site

import (
	"net/http"

	"github.com/kapmahc/axe"
	"github.com/kapmahc/sky/plugins/auth"
)

func (p *Plugin) indexAdminUsers(c *axe.Context) {
	var items []auth.User
	if err := p.Db.
		Order("last_sign_in_at DESC").Find(&items).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, items)
}
