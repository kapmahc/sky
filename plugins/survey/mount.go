package survey

import (
	"github.com/kapmahc/axe"
)

// Mount mount web points
func (p *Plugin) Mount(rt *axe.Router) {
	mg := axe.NewRouter()
	mg.POST("/forms/{id}/apply", p.postFormApply)
	mg.POST("/forms/{id}/cancel", p.postFormCancel)
	mg.GET("/forms/{id}/export", p.Jwt.MustAdminMiddleware, p.getFormExport)
	mg.Resources(
		"/forms",
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.indexForms},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.createForm},
		[]axe.HandlerFunc{p.showForm},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.updateForm},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.destroyForm},
	)

	rt.Group("/api/survey", mg)
}
