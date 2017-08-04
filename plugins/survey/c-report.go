package survey

import (
	"encoding/json"
	"net/http"

	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
)

func (p *Plugin) getFormReport(c *axe.Context) {
	var item Form
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Model(&item).Association("Fields").Find(&item.Fields).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Model(&item).Association("Records").Find(&item.Records).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	lang := c.Payload[i18n.LOCALE].(string)
	headers := []axe.H{
		axe.H{"name": "username", "label": p.I18n.T(lang, "forms.attributes.record.username")},
		axe.H{"name": "email", "label": p.I18n.T(lang, "forms.attributes.record.email")},
		axe.H{"name": "phone", "label": p.I18n.T(lang, "forms.attributes.record.phone")},
	}
	for _, f := range item.Fields {
		headers = append(headers, axe.H{"name": f.Name, "label": f.Label})
	}

	var rows []axe.H
	for _, r := range item.Records {
		row := axe.H{
			"username": r.Username,
			"email":    r.Email,
			"phone":    r.Phone,
		}
		val := make(map[string]interface{})
		if err := json.Unmarshal([]byte(r.Value), &val); err != nil {
			c.Abort(http.StatusInternalServerError, err)
			return
		}
		for _, f := range item.Fields {
			row[f.Name] = val[f.Name]
		}
		rows = append(rows, row)
	}

	c.JSON(http.StatusOK, axe.H{
		"headers": headers,
		"rows":    rows,
	})
}
