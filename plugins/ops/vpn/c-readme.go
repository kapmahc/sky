package vpn

import (
	"net/http"
	"path"

	"github.com/kapmahc/axe"
)

func (p *Plugin) getReadme(c *axe.Context) {
	c.TEXT(http.StatusOK, path.Join("ops", "vpn", "readme.md"), axe.H{})
}
