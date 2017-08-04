package reading

import (
	"net/http"

	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
	"github.com/kapmahc/sky/plugins/auth"
	"github.com/kapmahc/sky/web"
)

func (p *Plugin) getMyNotes(c *axe.Context) (interface{}, error) {

	user := c.MustGet(auth.CurrentUser).(*auth.User)
	isa := c.MustGet(auth.IsAdmin).(bool)
	var notes []Note
	qry := p.Db
	if !isa {
		qry = qry.Where("user_id = ?", user.ID)
	}
	if err := qry.Order("updated_at DESC").Find(&notes).Error; err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, notes)
	return nil
}

func (p *Plugin) indexNotes(c *axe.Context) (interface{}, error) {

	var total int64
	var pag *web.Pagination
	if err := p.Db.Model(&Note{}).Count(&total).Error; err != nil {
		return nil, err
	}

	pag = web.NewPagination(c.Request, total)
	var notes []Note
	if err := p.Db.
		Limit(pag.Limit()).Offset(pag.Offset()).
		Find(&notes).Error; err != nil {
		return nil, err
	}

	for _, it := range notes {
		pag.Items = append(pag.Items, it)
	}

	c.JSON(http.StatusOK, pag)
	return nil
}

type fmNoteNew struct {
	Type   string `json:"type" binding:"required,max=8"`
	Body   string `json:"body" binding:"required,max=2000"`
	BookID uint   `json:"bookId"`
}

func (p *Plugin) createNote(c *axe.Context) (interface{}, error) {

	user := c.MustGet(auth.CurrentUser).(*auth.User)

	var fm fmNoteNew
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	item := Note{
		Type:   fm.Type,
		Body:   fm.Body,
		BookID: fm.BookID,
		UserID: user.ID,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return nil, err
	}

	c.JSON(http.StatusOK, item)
	return nil
}

func (p *Plugin) showNote(c *axe.Context) (interface{}, error) {
	var item Note
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, item)
	return nil
}

type fmNoteEdit struct {
	Type string `json:"type" binding:"required,max=8"`
	Body string `json:"body" binding:"required,max=2000"`
}

func (p *Plugin) updateNote(c *axe.Context) (interface{}, error) {
	note := c.MustGet("item").(*Note)

	var fm fmNoteEdit
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	if err := p.Db.Model(note).
		Updates(map[string]interface{}{
			"body": fm.Body,
			"type": fm.Type,
		}).Error; err != nil {
		return nil, err
	}

	c.JSON(http.StatusOK, axe.H{})
	return nil
}

func (p *Plugin) destroyNote(c *axe.Context) (interface{}, error) {
	n := c.MustGet("item").(*Note)
	if err := p.Db.Delete(n).Error; err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, axe.H{})
	return nil
}

func (p *Plugin) canEditNote(c *axe.Context) (interface{}, error) {
	lng := c.Payload[i18n.LOCALE].(string)
	user := c.MustGet(auth.CurrentUser).(*auth.User)

	var n Note
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&n).Error; err != nil {
		return nil, err
	}
	if user.ID == n.UserID || c.MustGet(auth.IsAdmin).(bool) {
		c.Set("item", &n)
		return nil
	}
	return p.I18n.E(http.StatusForbidden, lng, "errors.forbidden")
}
