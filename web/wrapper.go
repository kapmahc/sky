package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

const (
	// LayoutApplication application layout
	LayoutApplication = "layouts/application/index"
	// LayoutDashboard dashboard layout
	LayoutDashboard = "layouts/dashboard/index"
)

// Wrapper wrapper
type Wrapper struct {
	Render *render.Render `inject:""`
}

// HTML render html
func (p *Wrapper) HTML(l, n string, f func(*gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := f(c); err != nil {
			log.Error(err)
			c.Set("error", err.Error())
		}
		p.Render.HTML(c.Writer, http.StatusOK, n, c.Keys, render.HTMLOptions{Layout: l})
	}
}

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
