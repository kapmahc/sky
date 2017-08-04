package survey

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
)

func (p *Plugin) _parseValues(f *Field) []interface{} {
	var items []interface{}
	for _, s := range strings.Split(f.Value, ";") {
		items = append(items, s)
	}
	return items
}

type fmApply struct {
	Records  []fmRecord `json:"records" binding:"required,max=255"`
	Username string     `json:"username" binding:"required,max=255"`
	Email    string     `json:"email" binding:"required,max=255"`
	Phone    string     `json:"phone" binding:"required,max=255"`
}

type fmRecord struct {
	Name  string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
}

func (p *Plugin) postFormApply(c *axe.Context) {
	var fm fmApply
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	var item Form
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	lng := c.Payload[i18n.LOCALE].(string)
	if item.Expire() {
		c.Abort(http.StatusInternalServerError, p.I18n.E(lng, "forms.errors.expired"))
		return
	}
	var count int
	if err := p.Db.Model(&Record{}).Where("form_id = ? AND (phone = ? OR email = ?)", item.ID, fm.Phone, fm.Email).Count(&count).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if count > 0 {
		c.Abort(http.StatusInternalServerError, p.I18n.E(lng, "forms.errors.already-apply"))
		return
	}

	values := make(map[string]interface{})
	for _, r := range fm.Records {
		values[r.Name] = r.Value
	}

	val, err := json.Marshal(values)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	record := Record{
		Email:    fm.Email,
		Phone:    fm.Phone,
		Username: fm.Username,
		Value:    string(val),
		FormID:   item.ID,
	}
	if err := p.Db.Create(&record).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	p._sendEmail(lng, &item, &record, actApply)

	c.JSON(http.StatusOK, axe.H{})
}
