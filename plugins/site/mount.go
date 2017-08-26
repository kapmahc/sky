package site

import "github.com/kapmahc/axe"

// Mount mount web points
func (p *Plugin) Mount(rt *axe.Router) {
	rt.FuncMap("t", func(lang, code string, args ...interface{}) string {
		return p.I18n.T(lang, code, args...)
	})

	// --------------------

	ag := axe.NewRouter()
	ag.Use(p.Jwt.MustAdminMiddleware)
	ag.GET("/locales", p.indexAdminLocale)
	ag.GET("/locales/{code}", p.showAdminLocale)
	ag.DELETE("/locales/{code}", p.destroyAdminLocale)
	ag.POST("/locales", p.postAdminLocales)
	ag.POST("/site/info", p.postAdminSiteInfo)
	ag.POST("/site/author", p.postAdminSiteAuthor)
	ag.GET("/seo", p.getAdminSeo)
	ag.POST("/seo", p.postAdminSeo)
	ag.GET("/smtp", p.getAdminSMTP)
	ag.POST("/smtp", p.postAdminSMTP)
	ag.GET("/status", p.getAdminStatus)
	ag.GET("/users", p.indexAdminUsers)
	ag.GET("/paypal", p.getPaypal)
	ag.POST("/paypal", p.postPaypal)
	rt.Group("/api/admin", ag)

	rt.POST("/api/install", p.postInstall)
	rt.GET("/api/site/info", p.getSiteInfo)
	rt.GET("/api/locales/{lang}", p.getLocales)
	rt.Resources(
		"/api/cards",
		[]axe.HandlerFunc{p.indexCards},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.createCard},
		[]axe.HandlerFunc{p.showCard},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.updateCard},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.destroyCard},
	)
	rt.Resources(
		"/api/links",
		[]axe.HandlerFunc{p.indexLinks},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.createLink},
		[]axe.HandlerFunc{p.showLink},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.updateLink},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.destroyLink},
	)
	rt.Resources(
		"/api/friend-links",
		[]axe.HandlerFunc{p.indexFriendLinks},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.createFriendLink},
		[]axe.HandlerFunc{p.showFriendLink},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.updateFriendLink},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.destroyFriendLink},
	)
	rt.Resources(
		"/api/leave-words",
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.indexLeaveWords},
		[]axe.HandlerFunc{p.createLeaveWord},
		nil,
		nil,
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.destroyLeaveWord},
	)

	// -----------------
	rt.GET("/", p.getHome)
}
