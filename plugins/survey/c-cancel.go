package survey

import (
	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
)

type fmCancel struct {
	Who string `json:"who" binding:"required,max=255"`
}

func (p *Plugin) postFormCancel(c *axe.Context) (interface{}, error) {
	var fm fmCancel
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	lng := c.Payload[i18n.LOCALE].(string)
	var item Form
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		return nil, err
	}

	if item.Expire() {
		return nil, p.I18n.E(lng, "forms.errors.expired")
	}
	var record Record
	if err := p.Db.Where("form_id = ? AND (phone = ? OR email = ?)", item.ID, fm.Who, fm.Who).First(&record).Error; err != nil {
		return nil, err
	}

	if err := p.Db.Delete(&record).Error; err != nil {
		return nil, err
	}
	p._sendEmail(lng, &item, &record, actCancel)

	return axe.H{}, nil
}
