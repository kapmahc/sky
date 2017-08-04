package auth

import (
	"net/http"

	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
)

func (p *Plugin) createAttachment(c *axe.Context) (interface{}, error) {
	user := c.Payload[CurrentUser].(*User)

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return nil, err
	}

	url, size, err := p.Uploader.Save(header)
	if err != nil {
		return nil, err
	}

	// http://golang.org/pkg/net/http/#DetectContentType
	buf := make([]byte, 512)
	file.Seek(0, 0)
	if _, err = file.Read(buf); err != nil {
		return nil, err
	}

	a := Attachment{
		Title:        header.Filename,
		URL:          url,
		UserID:       user.ID,
		MediaType:    http.DetectContentType(buf),
		Length:       size / 1024,
		ResourceType: DefaultResourceType, //fm.Type,
		ResourceID:   DefaultResourceID,   //fm.ID,
	}
	if err := p.Db.Create(&a).Error; err != nil {
		return nil, err
	}
	return axe.H{
		"url":    a.URL,
		"uid":    a.ID,
		"status": "success",
	}, nil
}

type fmAttachmentEdit struct {
	Title string `json:"title" binding:"required,max=255"`
}

func (p *Plugin) updateAttachment(c *axe.Context) (interface{}, error) {
	a, err := p.canEditAttachment(c)
	if err != nil {
		return nil, err
	}
	var fm fmAttachmentEdit
	if err = c.Bind(&fm); err != nil {
		return nil, err
	}
	if err = p.Db.Model(a).Update("title", fm.Title).Error; err != nil {
		return nil, err
	}

	return a, nil
}

func (p *Plugin) destroyAttachment(c *axe.Context) (interface{}, error) {
	a, err := p.canEditAttachment(c)
	if err != nil {
		return nil, err
	}
	if err := p.Db.Delete(a).Error; err != nil {
		return nil, err
	}
	if err := p.Uploader.Remove(a.URL); err != nil {
		return nil, err
	}
	return a, nil
}

func (p *Plugin) showAttachment(c *axe.Context) (interface{}, error) {
	var a Attachment
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (p *Plugin) indexAttachments(c *axe.Context) (interface{}, error) {
	user := c.Payload[CurrentUser].(*User)
	isa := c.Payload[IsAdmin].(bool)
	var items []Attachment
	qry := p.Db
	if !isa {
		qry = qry.Where("user_id = ?", user.ID)
	}
	if err := qry.Order("updated_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (p *Plugin) canEditAttachment(c *axe.Context) (*Attachment, error) {
	user := c.Payload[CurrentUser].(*User)
	lng := c.Payload[i18n.LOCALE].(string)

	var a Attachment
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&a).Error; err != nil {
		return nil, err
	}

	if user.ID == a.UserID || c.Payload[IsAdmin].(bool) {
		return &a, nil
	}

	return nil, p.I18n.E(lng, "errors.not-allow")
}
