package reading

import (
	"github.com/kapmahc/axe"
)

// Mount mount web points
func (p *Plugin) Mount(rt *axe.Router) {
	ug := axe.NewRouter()
	ug.GET(`/pages/{id}/{href:[/.\w]+}`, p.showPage)
	rt.Group("/reading", ug)

	mg := axe.NewRouter()
	mg.Resources(
		"/books",
		[]axe.HandlerFunc{p.indexBooks},
		nil,
		[]axe.HandlerFunc{p.showBook},
		nil,
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.destroyBook},
	)
	mg.Resources(
		"/notes",
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, p.indexNotes},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, p.createNote},
		[]axe.HandlerFunc{p.showNote},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, p.updateNote},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, p.destroyNote},
	)
	rt.Group("/api/reading", mg)

}
