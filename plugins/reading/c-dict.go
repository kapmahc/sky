package reading

import (
	"net/http"

	"github.com/kapmahc/axe"
)

type fmDict struct {
	Keywords string `json:"keywords" binding:"required,max=255"`
}

func (p *Plugin) postDict(c *axe.Context) {

	var fm fmDict
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	rst := axe.H{}
	for _, dic := range dictionaries {
		for _, sen := range dic.Translate(fm.Keywords) {
			var items []axe.H
			for _, pat := range sen.Parts {
				items = append(items, axe.H{"type": pat.Type, "body": string(pat.Data)})
			}
			rst[dic.GetBookName()] = items
		}
	}

	c.JSON(http.StatusOK, rst)
}
