package survey

import (
	"net/http"
	"time"

	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
	"github.com/kapmahc/sky/web"
)

func (p *Plugin) indexForms(c *axe.Context) {
	var items []Form
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, items)
}

func (p *Plugin) createForm(c *axe.Context) {
	var fm fmForm
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	deadline, err := time.Parse(time.RFC3339, fm.Deadline)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	item := Form{
		Title:    fm.Title,
		Deadline: deadline,
		Media: web.Media{
			Body: fm.Body,
			Type: fm.Type,
		},
	}
	if err := p.Db.Create(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	if err := p.Db.Model(&item).Association("Fields").Append(fm.Parse(item.ID)).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (p *Plugin) showForm(c *axe.Context) {
	var item Form
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Model(&item).Association("Fields").Find(&item.Fields).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

type fmField struct {
	Name  string `json:"name" binding:"required"`
	Label string `json:"label" binding:"required"`
	Body  string `json:"body"`
	Value string `json:"value"`
	Type  string `json:"type" binding:"required"`
}

type fmForm struct {
	Title    string    `json:"title" binding:"required,max=255"`
	Deadline string    `json:"deadline" binding:"required"`
	Body     string    `json:"body" binding:"required"`
	Type     string    `json:"type" binding:"required,max=8"`
	Fields   []fmField `json:"fields"`
}

func (p *fmForm) Parse(id uint) []Field {
	var fields []Field
	for i, f := range p.Fields {
		fields = append(
			fields, Field{
				FormID:    id,
				SortOrder: i,
				Name:      f.Name,
				Label:     f.Label,
				Value:     f.Value,
				Media: web.Media{
					Body: f.Body,
					Type: f.Type,
				},
			})
	}
	return fields
}

func (p *Plugin) updateForm(c *axe.Context) {
	var item Form
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	var fm fmForm
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	deadline, err := time.Parse(time.RFC3339, fm.Deadline)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	if err := p.Db.Model(&item).Updates(map[string]interface{}{
		"title":    fm.Title,
		"type":     fm.Type,
		"body":     fm.Body,
		"deadline": deadline,
	}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	lang := c.Payload[i18n.LOCALE].(string)
	if p.Db.Model(&item).Association("Records").Count() > 0 {
		c.Abort(http.StatusInternalServerError, p.I18n.E(lang, "errors.in-use"))
		return
	}

	if err := p.Db.Where("form_id = ?", item.ID).Delete(&Field{}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	if err := p.Db.Model(&item).Association("Fields").Append(fm.Parse(item.ID)).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (p *Plugin) destroyForm(c *axe.Context) {
	id := c.Params["id"]
	if err := p.Db.Where("form_id = ?", id).Delete(Field{}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Where("form_id = ?", id).Delete(Record{}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Where("id = ?", id).Delete(Form{}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, axe.H{})
}
