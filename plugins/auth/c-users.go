package auth

import (
	"net/http"

	"github.com/kapmahc/sky/web"
)

func (p *Plugin) getUsersSignIn(wrt http.ResponseWriter, req *http.Request) {
	p.Render.JSON(wrt, http.StatusOK, web.H{"ok": true})
}
