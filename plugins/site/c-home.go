package site

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Plugin) getHome(c *gin.Context) error {
	return nil
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
