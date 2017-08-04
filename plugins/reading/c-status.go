package reading

import (
	"net/http"

	"github.com/kapmahc/axe"
)

func (p *Plugin) getStatus(c *axe.Context) {
	data := axe.H{}
	var bc int
	if err := p.Db.Model(&Book{}).Count(&bc).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	data["book"] = axe.H{
		"count": bc,
	}

	dict := axe.H{}
	for _, dic := range dictionaries {
		dict[dic.GetBookName()] = dic.GetWordCount()
	}
	data["dict"] = dict

	c.JSON(http.StatusOK, data)

}
