package mail

import (
	"github.com/kapmahc/axe"
)

// Mount mount web points
func (p *Plugin) Mount(rt *axe.Router) {
	mg := axe.NewRouter()
	mg.GET("/readme", p.Jwt.MustAdminMiddleware, p.getReadme)
	mg.Resources(
		"/aliases",
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.indexAliases},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.createAlias},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.showAlias},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.updateAlias},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.destroyAlias},
	)
	mg.Resources(
		"/domains",
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.indexDomains},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.createDomain},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.showDomain},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.updateDomain},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.destroyDomain},
	)
	mg.POST("/users/reset-password", p.Jwt.MustAdminMiddleware, p.postResetUserPassword)
	mg.POST("/users/change-password", p.postChangeUserPassword)
	mg.Resources(
		"/users",
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.indexUsers},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.createUser},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.showUser},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.updateUser},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.destroyUser},
	)

	rt.Group("/api/ops/mail", mg)
}
