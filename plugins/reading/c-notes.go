package reading

import (
	"net/http"

	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
	"github.com/kapmahc/sky/plugins/auth"
)

func (p *Plugin) indexNotes(c *axe.Context) {

	user := c.Payload[auth.CurrentUser].(*auth.User)
	isa := c.Payload[auth.IsAdmin].(bool)
	var notes []Note
	qry := p.Db
	if !isa {
		qry = qry.Where("user_id = ?", user.ID)
	}
	if err := qry.Order("updated_at DESC").Find(&notes).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, notes)
}

type fmNoteNew struct {
	Type   string `json:"type" validate:"required,max=8"`
	Body   string `json:"body" validate:"required,max=2000"`
	BookID uint   `json:"bookId"`
}

func (p *Plugin) createNote(c *axe.Context) {

	user := c.Payload[auth.CurrentUser].(*auth.User)

	var fm fmNoteNew
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	item := Note{
		Type:   fm.Type,
		Body:   fm.Body,
		BookID: fm.BookID,
		UserID: user.ID,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (p *Plugin) showNote(c *axe.Context) {
	var item Note
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

type fmNoteEdit struct {
	Type string `json:"type" validate:"required,max=8"`
	Body string `json:"body" validate:"required,max=2000"`
}

func (p *Plugin) updateNote(c *axe.Context) {
	note, err := p.canEditNote(c)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	var fm fmNoteEdit
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	if err := p.Db.Model(note).
		Updates(map[string]interface{}{
			"body": fm.Body,
			"type": fm.Type,
		}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, note)
}

func (p *Plugin) destroyNote(c *axe.Context) {
	n, err := p.canEditNote(c)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Delete(n).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, n)
}

func (p *Plugin) canEditNote(c *axe.Context) (*Note, error) {
	lng := c.Payload[i18n.LOCALE].(string)
	user := c.Payload[auth.CurrentUser].(*auth.User)

	var n Note
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&n).Error; err != nil {
		return nil, err
	}
	if user.ID == n.UserID || c.Payload[auth.IsAdmin].(bool) {
		return &n, nil
	}
	return nil, p.I18n.E(lng, "errors.forbidden")
}
