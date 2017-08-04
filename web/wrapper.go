package web

import (
	"net/http"

	"github.com/kapmahc/axe"
)

// JSON wrapper json
func JSON(f func(*axe.Context) (interface{}, error)) axe.HandlerFunc {
	return func(c *axe.Context) {
		v, e := f(c)
		if e != nil {
			c.Abort(http.StatusInternalServerError, e)
			return
		}
		c.JSON(http.StatusOK, v)
	}
}
