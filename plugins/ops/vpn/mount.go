package vpn

import (
	"github.com/kapmahc/axe"
)

// Mount mount web points
func (p *Plugin) Mount(rt *axe.Router) {
	mg := axe.NewRouter()
	mg.GET("/readme", p.Jwt.MustAdminMiddleware, p.getReadme)
	mg.GET("/logs", p.Jwt.MustSignInMiddleware, p.indexLogs)

	mg.PATCH("/users/auth", p.tokenMiddleware, p.apiAuth)
	mg.PATCH("/users/connect", p.tokenMiddleware, p.apiConnect)
	mg.PATCH("/users/disconnect", p.tokenMiddleware, p.apiDisconnect)

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

	rt.Group("/api/ops/vpn", mg)

}
