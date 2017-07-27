package site

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapmahc/sky/web"
)

func (p *Plugin) getHome(c *gin.Context) {
	// TODO
	p.Render.HTML(c.Writer, http.StatusOK, "site/home", gin.H{}, web.LayoutApplication)
}

func (p *Plugin) getDonates(c *gin.Context) error {
	data := gin.H{}
	var paypal map[string]interface{}
	if err := p.Settings.Get("site.paypal", &paypal); err == nil {
		data["paypal"] = paypal["donate"]
	}
	c.JSON(http.StatusOK, data)
	return nil
}
