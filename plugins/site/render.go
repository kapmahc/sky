package site

import (
	"html/template"
	"path"

	"github.com/kapmahc/sky/web"
	"github.com/unrolled/render"
)

func (p *Plugin) openRender(theme string) *render.Render {
	return render.New(render.Options{
		Directory:     path.Join("themes", theme, "views"),
		Extensions:    []string{".html"},
		Funcs:         []template.FuncMap{},
		IsDevelopment: !web.IsProduction(),
	})
}
