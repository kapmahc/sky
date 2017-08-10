package forum

import (
	"net/http"

	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/i18n"
	"github.com/kapmahc/sky/plugins/auth"
)

func (p *Plugin) indexArticles(c *axe.Context) {
	user := c.Payload[auth.CurrentUser].(*auth.User)
	isa := c.Payload[auth.IsAdmin].(bool)
	var articles []Article
	qry := p.Db.Select([]string{"title", "updated_at", "id"})
	if !isa {
		qry = qry.Where("user_id = ?", user.ID)
	}
	if err := qry.Order("updated_at DESC").Find(&articles).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, articles)
}

type fmArticle struct {
	Title   string   `json:"title" validate:"required,max=255"`
	Summary string   `json:"summary" validate:"required,max=500"`
	Type    string   `json:"type" validate:"required,max=8"`
	Body    string   `json:"body" validate:"required,max=2000"`
	Tags    []string `json:"tags"`
}

func (p *Plugin) createArticle(c *axe.Context) {
	user := c.Payload[auth.CurrentUser].(*auth.User)
	var fm fmArticle
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	var tags []Tag
	for _, it := range fm.Tags {
		var t Tag
		if err := p.Db.Select([]string{"id"}).Where("id = ?", it).First(&t).Error; err == nil {
			tags = append(tags, t)
		} else {
			c.Abort(http.StatusInternalServerError, err)
			return
		}
	}
	a := Article{
		Title:   fm.Title,
		Summary: fm.Summary,
		Body:    fm.Body,
		Type:    fm.Type,
		UserID:  user.ID,
	}

	if err := p.Db.Create(&a).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Model(&a).Association("Tags").Append(tags).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &a)
}

func (p *Plugin) showArticle(c *axe.Context) {
	var a Article
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&a).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Model(&a).Related(&a.Comments).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Model(&a).Association("Tags").Find(&a.Tags).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, a)
}

func (p *Plugin) updateArticle(c *axe.Context) {
	a, err := p.canEditArticle(c)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	var fm fmArticle
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	var tags []Tag
	for _, it := range fm.Tags {
		var t Tag
		if err := p.Db.Select([]string{"id"}).Where("id = ?", it).First(&t).Error; err == nil {
			tags = append(tags, t)
		} else {
			c.Abort(http.StatusInternalServerError, err)
			return
		}
	}

	if err := p.Db.Model(a).Updates(map[string]interface{}{
		"title":   fm.Title,
		"summary": fm.Summary,
		"body":    fm.Body,
		"type":    fm.Type,
	}).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	if err := p.Db.Model(a).Association("Tags").Replace(tags).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, a)
}

func (p *Plugin) destroyArticle(c *axe.Context) {
	a, err := p.canEditArticle(c)
	if err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Model(a).Association("Tags").Clear().Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Delete(a).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, a)
}

func (p *Plugin) canEditArticle(c *axe.Context) (*Article, error) {
	lng := c.Payload[i18n.LOCALE].(string)
	user := c.Payload[auth.CurrentUser].(*auth.User)

	var a Article
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&a).Error; err != nil {
		return nil, err
	}

	if user.ID == a.UserID || c.Payload[auth.IsAdmin].(bool) {
		return &a, nil
	}
	return nil, p.I18n.E(lng, "errors.forbidden")
}
