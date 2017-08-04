package auth

import (
	"github.com/kapmahc/axe"
)

// Mount mount web points
func (p *Plugin) Mount(rt *axe.Router) {
	ug := axe.NewRouter()
	ug.POST("/sign-in", p.postUsersSignIn)
	ug.POST("/sign-up", p.postUsersSignUp)
	ug.POST("/confirm", p.postUsersConfirm)
	ug.GET("/confirm/{token}", p.getUsersConfirm)
	ug.POST("/unlock", p.postUsersUnlock)
	ug.GET("/unlock/{token}", p.getUsersUnlock)
	ug.POST("/forgot-password", p.postUsersForgotPassword)
	ug.POST("/reset-password", p.postUsersResetPassword)
	ug.GET("/logs", p.Jwt.MustSignInMiddleware, p.getUsersLogs)
	ug.GET("/info", p.Jwt.MustSignInMiddleware, p.getUsersInfo)
	ug.POST("/info", p.Jwt.MustSignInMiddleware, p.postUsersInfo)
	ug.POST("/change-password", p.Jwt.MustSignInMiddleware, p.postUsersChangePassword)
	ug.DELETE("/sign-out", p.Jwt.MustSignInMiddleware, p.deleteUsersSignOut)
	rt.Group("/api/users", ug)

	rt.Resources(
		"/api/attachments",
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, p.indexAttachments},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, p.createAttachment},
		[]axe.HandlerFunc{p.showAttachment},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, p.updateAttachment},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, p.destroyAttachment},
	)
}
