package site

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/sky/plugins/auth"
	"github.com/kapmahc/sky/web"
)

// Mount mount web points
func (p *Plugin) Mount(rt *gin.Engine) {
	rt.GET("/", p.Wrapper.HTML(web.LayoutApplication, "site/home", p.getHome))
	rt.GET("/install", web.Wrap(p.mustDatabaseEmpty), p.Wrapper.HTML(auth.LayoutNonSignIn, "site/install", p.getInstall))
	// --------------

	rt.POST("/install", web.Wrap(p.mustDatabaseEmpty), web.Wrap(p.postInstall))
	rt.GET("/donates", web.Wrap(p.getDonates))

	ag := rt.Group("/admin", web.Wrap(p.Jwt.MustAdminMiddleware))

	ag.GET("/users", web.Wrap(p.indexAdminUsers))

	ag.GET("/locales", web.Wrap(p.indexAdminLocales))
	ag.POST("/locales", web.Wrap(p.postAdminLocales))
	ag.GET("/locales/:code", web.Wrap(p.showAdminLocale))
	ag.DELETE("/locales/:code", web.Wrap(p.destroyAdminLocale))

	asg := ag.Group("/site")
	asg.GET("/status", web.Wrap(p.getAdminSiteStatus))
	asg.POST("/info", web.Wrap(p.postAdminSiteInfo))
	asg.POST("/author", web.Wrap(p.postAdminSiteAuthor))
	asg.GET("/seo", web.Wrap(p.getAdminSiteSeo))
	asg.POST("/seo", web.Wrap(p.postAdminSiteSeo))
	asg.GET("/smtp", web.Wrap(p.getAdminSiteSMTP))
	asg.POST("/smtp", web.Wrap(p.postAdminSiteSMTP))
	asg.GET("/paypal", web.Wrap(p.getAdminSitePaypal))
	asg.POST("/paypal", web.Wrap(p.postAdminSitePaypal))

	rt.GET("/leave-words", web.Wrap(p.Jwt.MustAdminMiddleware), web.Wrap(p.indexLeaveWords))
	rt.POST("/leave-words", web.Wrap(p.createLeaveWord))
	rt.DELETE("/leave-words/:id", web.Wrap(p.Jwt.MustAdminMiddleware), web.Wrap(p.destroyLeaveWord))

	rt.GET("/friend-links", web.Wrap(p.indexFriendLinks))
	rt.GET("/friend-links/:id", web.Wrap(p.showFriendLink))
	rt.POST("/friend-links", web.Wrap(p.Jwt.MustAdminMiddleware), web.Wrap(p.createFriendLink))
	rt.POST("/friend-links/:id", web.Wrap(p.Jwt.MustAdminMiddleware), web.Wrap(p.updateFriendLink))
	rt.DELETE("/friend-links/:id", web.Wrap(p.Jwt.MustAdminMiddleware), web.Wrap(p.destroyFriendLink))

	rt.GET("/links", web.Wrap(p.indexLinks))
	rt.GET("/links/:id", web.Wrap(p.showLink))
	rt.POST("/links", web.Wrap(p.Jwt.MustAdminMiddleware), web.Wrap(p.createLink))
	rt.POST("/links/:id", web.Wrap(p.Jwt.MustAdminMiddleware), web.Wrap(p.updateLink))
	rt.DELETE("/links/:id", web.Wrap(p.Jwt.MustAdminMiddleware), web.Wrap(p.destroyLink))

	rt.GET("/cards", web.Wrap(p.indexCards))
	rt.GET("/cards/:id", web.Wrap(p.showCard))
	rt.POST("/cards", web.Wrap(p.Jwt.MustAdminMiddleware), web.Wrap(p.createCard))
	rt.POST("/cards/:id", web.Wrap(p.Jwt.MustAdminMiddleware), web.Wrap(p.updateCard))
	rt.DELETE("/cards/:id", web.Wrap(p.Jwt.MustAdminMiddleware), web.Wrap(p.destroyCard))

}
