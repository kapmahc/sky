package auth

import (
	"github.com/kapmahc/axe"
	"github.com/kapmahc/sky/web"
)

// Mount mount web points
func (p *Plugin) Mount(rt *axe.Router) {
	ug := axe.NewRouter()
	ug.POST("/sign-in", web.JSON(p.postUsersSignIn))
	ug.POST("/sign-up", web.JSON(p.postUsersSignUp))
	ug.POST("/confirm", web.JSON(p.postUsersConfirm))
	ug.GET("/confirm/{token}", web.JSON(p.getUsersConfirm))
	ug.POST("/unlock", web.JSON(p.postUsersUnlock))
	ug.GET("/unlock/{token}", web.JSON(p.getUsersUnlock))
	ug.POST("/forgot-password", web.JSON(p.postUsersForgotPassword))
	ug.POST("/reset-password", web.JSON(p.postUsersResetPassword))
	ug.GET("/logs", p.Jwt.MustSignInMiddleware, web.JSON(p.getUsersLogs))
	ug.GET("/info", p.Jwt.MustSignInMiddleware, web.JSON(p.getUsersInfo))
	ug.POST("/info", p.Jwt.MustSignInMiddleware, web.JSON(p.postUsersInfo))
	ug.POST("/change-password", p.Jwt.MustSignInMiddleware, web.JSON(p.postUsersChangePassword))
	ug.DELETE("/sign-out", p.Jwt.MustSignInMiddleware, web.JSON(p.deleteUsersSignOut))
	rt.Group("/api/users", ug)

	rt.Resources(
		"/api/attachments",
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, web.JSON(p.indexAttachments)},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, web.JSON(p.createAttachment)},
		[]axe.HandlerFunc{web.JSON(p.showAttachment)},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, web.JSON(p.updateAttachment)},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, web.JSON(p.destroyAttachment)},
	)
}
