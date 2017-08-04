package forum

import (
	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
	"github.com/kapmahc/sky/plugins/auth"
)

func (p *Plugin) indexComments(c *axe.Context) (interface{}, error) {

	user := c.Payload[auth.CurrentUser].(*auth.User)
	isa := c.Payload[auth.IsAdmin].(bool)
	var comments []Comment
	qry := p.Db.Select([]string{"body", "article_id", "updated_at", "id"})
	if !isa {
		qry = qry.Where("user_id = ?", user.ID)
	}
	if err := qry.Order("updated_at DESC").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

type fmCommentAdd struct {
	Body      string `json:"body" binding:"required,max=800"`
	Type      string `json:"type" binding:"required,max=8"`
	ArticleID uint   `json:"articleId" binding:"required"`
}

func (p *Plugin) createComment(c *axe.Context) (interface{}, error) {

	user := c.Payload[auth.CurrentUser].(*auth.User)

	var fm fmCommentAdd
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	cm := Comment{
		Body:      fm.Body,
		Type:      fm.Type,
		ArticleID: fm.ArticleID,
		UserID:    user.ID,
	}

	if err := p.Db.Create(&cm).Error; err != nil {
		return nil, err
	}

	return cm, nil
}

func (p *Plugin) showComment(c *axe.Context) (interface{}, error) {
	var item Comment
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

type fmCommentEdit struct {
	Body string `json:"body" binding:"required,max=800"`
	Type string `json:"type" binding:"required,max=8"`
}

func (p *Plugin) updateComment(c *axe.Context) (interface{}, error) {
	cm, err := p.canEditComment(c)
	if err != nil {
		return nil, err
	}

	var fm fmCommentEdit
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	if err := p.Db.Model(cm).Updates(map[string]interface{}{
		"body": fm.Body,
		"type": fm.Type,
	}).Error; err != nil {
		return nil, err
	}

	return cm, nil
}

func (p *Plugin) destroyComment(c *axe.Context) (interface{}, error) {
	cm, err := p.canEditComment(c)
	if err != nil {
		return nil, err
	}
	if err := p.Db.Delete(cm).Error; err != nil {
		return nil, err
	}
	return cm, nil
}

func (p *Plugin) canEditComment(c *axe.Context) (*Comment, error) {
	lng := c.Payload[i18n.LOCALE].(string)
	user := c.Payload[auth.CurrentUser].(*auth.User)

	var o Comment
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&o).Error; err != nil {
		return nil, nil
	}

	if user.ID == o.UserID || c.Payload[auth.IsAdmin].(bool) {
		return &o, nil
	}
	return nil, p.I18n.E(lng, "errors.forbidden")
}
