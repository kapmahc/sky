package survey

import (
	"net/http"

	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
)

type fmCancel struct {
	Who string `json:"who" binding:"required,max=255"`
}

func (p *Plugin) postFormCancel(c *axe.Context) {
	var fm fmCancel
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	lng := c.Payload[i18n.LOCALE].(string)
	var item Form
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	if item.Expire() {
		c.Abort(http.StatusInternalServerError, p.I18n.E(lng, "survey.errors.expired"))
		return
	}
	var record Record
	if err := p.Db.Where("form_id = ? AND (phone = ? OR email = ?)", item.ID, fm.Who, fm.Who).First(&record).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	if err := p.Db.Delete(&record).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	p._sendEmail(lng, &item, &record, actCancel)

	c.JSON(http.StatusOK, axe.H{})
}
