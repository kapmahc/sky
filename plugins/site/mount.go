package site

import "github.com/kapmahc/axe"

// Mount mount web points
func (p *Plugin) Mount(rt *axe.Router) {
	rt.FuncMap("t", func(lang, code string, args ...interface{}) string {
		return p.I18n.T(lang, code, args...)
	})
}
