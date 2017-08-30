package survey

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kapmahc/axe"
)

func (p *Plugin) _parseValues(f *Field) []interface{} {
	var items []interface{}
	for _, s := range strings.Split(f.Value, ";") {
		items = append(items, s)
	}
	return items
}

type fmApply struct {
	Records  []fmRecord `json:"records" validate:"required,max=255"`
	Username string     `json:"username" validate:"required,max=255"`
	Email    string     `json:"email" validate:"required,max=255"`
	Phone    string     `json:"phone" validate:"required,max=255"`
}

type fmRecord struct {
	Name  string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
}

func (p *Plugin) postFormApply(c *axe.Context) {
	item, err := p.parseToken(c, actApply)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	var fm fmApply
	if err = c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
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
		Value:  string(val),
		FormID: item.ID,
	}
	if err := p.Db.Create(&record).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, axe.H{})
}
