package forum

import (
	"github.com/kapmahc/axe"
	"github.com/kapmahc/sky/web"
)

// Mount mount web points
func (p *Plugin) Mount(rt *axe.Router) {
	fg := axe.NewRouter()
	fg.Resources(
		"/articles",
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, web.JSON(p.indexArticles)},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, web.JSON(p.createArticle)},
		[]axe.HandlerFunc{web.JSON(p.showArticle)},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, web.JSON(p.updateArticle)},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, web.JSON(p.destroyArticle)},
	)
	fg.Resources(
		"/comments",
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, web.JSON(p.indexComments)},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, web.JSON(p.createComment)},
		[]axe.HandlerFunc{web.JSON(p.showComment)},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, web.JSON(p.updateComment)},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, web.JSON(p.destroyComment)},
	)
	fg.Resources(
		"/tags",
		[]axe.HandlerFunc{web.JSON(p.indexTags)},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, web.JSON(p.createTag)},
		[]axe.HandlerFunc{web.JSON(p.showTag)},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, web.JSON(p.updateTag)},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, web.JSON(p.destroyTag)},
	)
	rt.Group("/api/forums", fg)
}
