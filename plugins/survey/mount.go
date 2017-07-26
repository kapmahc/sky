package survey

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/sky/web"
)

// Mount mount web points
func (p *Plugin) Mount(rt *gin.Engine) {
	ag := rt.Group("/survey", web.Wrap(p.Jwt.MustAdminMiddleware))
	ag.POST("/", web.Wrap(p.createForm))
	ag.POST("/:id", web.Wrap(p.updateForm))
	ag.DELETE("/:id", web.Wrap(p.destroyForm))
	ag.GET("/:id/export", web.Wrap(p.getFormExport))
	ag.GET("/:id/report", web.Wrap(p.getFormReport))

	ng := rt.Group("/survey")
	ng.GET("/", web.Wrap(p.indexForms))
	ng.GET("/:id", web.Wrap(p.showForm))
	ng.POST("/:id/apply", web.Wrap(p.postFormApply))
	ng.POST("/:id/cancel", web.Wrap(p.postFormCancel))

}
