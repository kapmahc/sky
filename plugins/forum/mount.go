package forum

import (
	"github.com/kapmahc/axe"
)

// Mount mount web points
func (p *Plugin) Mount(rt *axe.Router) {
	fg := axe.NewRouter()
	fg.Resources(
		"/articles",
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, p.indexArticles},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, p.createArticle},
		[]axe.HandlerFunc{p.showArticle},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, p.updateArticle},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, p.destroyArticle},
	)
	fg.Resources(
		"/comments",
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, p.indexComments},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, p.createComment},
		[]axe.HandlerFunc{p.showComment},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, p.updateComment},
		[]axe.HandlerFunc{p.Jwt.MustSignInMiddleware, p.destroyComment},
	)
	fg.Resources(
		"/tags",
		[]axe.HandlerFunc{p.indexTags},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.createTag},
		[]axe.HandlerFunc{p.showTag},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.updateTag},
		[]axe.HandlerFunc{p.Jwt.MustAdminMiddleware, p.destroyTag},
	)
	rt.Group("/api/forum", fg)
}
