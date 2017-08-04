package vpn

import (
	"github.com/kapmahc/axe"
	"github.com/kapmahc/sky/web"
)

func (p *Plugin) indexLogs(c *axe.Context) (interface{}, error) {
	var total int64
	if err := p.Db.Model(&Log{}).Count(&total).Error; err != nil {
		return nil, err
	}
	pag := web.NewPagination(c.Request, total)

	var items []Log
	if err := p.Db.
		Limit(pag.Limit()).Offset(pag.Offset()).
		Find(&items).Error; err != nil {
		return nil, err
	}
	for _, b := range items {
		pag.Items = append(pag.Items, b)
	}
	return pag, nil
}
