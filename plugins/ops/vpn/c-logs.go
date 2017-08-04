package vpn

import (
	"net/http"

	"github.com/kapmahc/axe"
	"github.com/kapmahc/sky/web"
)

func (p *Plugin) indexLogs(c *axe.Context) {
	var total int64
	if err := p.Db.Model(&Log{}).Count(&total).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	pag := web.NewPagination(c.Request, total)

	var items []Log
	if err := p.Db.
		Limit(pag.Limit()).Offset(pag.Offset()).
		Find(&items).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	for _, b := range items {
		pag.Items = append(pag.Items, b)
	}

	c.JSON(http.StatusOK, pag)
}
