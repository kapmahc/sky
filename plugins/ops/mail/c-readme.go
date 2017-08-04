package mail

import (
	"net/http"
	"path"

	"github.com/kapmahc/axe"
)

func (p *Plugin) getReadme(c *axe.Context) {
	c.TEXT(http.StatusOK, path.Join("ops", "mail", "readme.md"), axe.H{})
}
