package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

var (
	// LayoutApplication application layout
	LayoutApplication = render.HTMLOptions{Layout: "layouts/application/index"}
	// LayoutDashboard dashboard layout
	LayoutDashboard = render.HTMLOptions{Layout: "layouts/dashboard/index"}
)

// Wrap wrap handler
func Wrap(f func(*gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if e := f(c); e != nil {
			log.Error(e)
			s := http.StatusInternalServerError
			if he, ok := e.(*HTTPError); ok {
				s = he.Status
			}
			c.String(s, e.Error())
			c.Abort()
		}
	}
}
