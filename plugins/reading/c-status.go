package reading

import (
	"net/http"

	"github.com/kapmahc/axe"
)

func (p *Plugin) getStatus(c *axe.Context) (interface{}, error) {
	data := axe.H{}
	var bc int
	if err := p.Db.Model(&Book{}).Count(&bc).Error; err != nil {
		return nil, err
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
	return nil

}
