package site

// https://developer.paypal.com/docs/directory/

import (
	"net/http"

	"github.com/kapmahc/axe"
)

type fmPaypal struct {
	Donate string `json:"donate"`
}

func (p *Plugin) getPaypal(c *axe.Context) {
	paypal := make(map[string]interface{})
	if err := p.Settings.Get("site.paypal", &paypal); err != nil {
		paypal["donate"] = ""
	}

	c.JSON(http.StatusOK, paypal)
}

func (p *Plugin) postPaypal(c *axe.Context) {
	var fm fmPaypal
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	paypal := map[string]interface{}{
		"donate": fm.Donate,
	}
	if err := p.Settings.Set("site.paypal", paypal, true); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, axe.H{})
}
