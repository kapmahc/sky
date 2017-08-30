package survey

import (
	"github.com/kapmahc/axe"
)

// Mount mount web points
func (p *Plugin) Mount(rt *axe.Router) {
	mg := axe.NewRouter()
	mg.POST("/apply", p.postFormApply)
	mg.POST("/cancel", p.postFormCancel)
	mg.GET("/export", p.getFormExport)
	mg.Resources(
		"/forms",
		[]axe.HandlerFunc{p.indexForms},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.createForm},
		[]axe.HandlerFunc{p.showForm},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.updateForm},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.destroyForm},
	)

	rt.Group("/api/survey", mg)
}
