package site

import (
	"net/http"

	"github.com/kapmahc/axe"
	"github.com/kapmahc/sky/web"
)

func (p *Plugin) getHome(c *axe.Context) {
	c.HTML(http.StatusOK, web.LayoutApplication, "site/home", axe.H{})
}
