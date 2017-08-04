package forum

import (
	"net/http"

	"github.com/kapmahc/axe"
)

func (p *Plugin) indexTags(c *axe.Context) {
	var tags []Tag
	if err := p.Db.Find(&tags).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, tags)
}

type fmTag struct {
	Name string `json:"name" binding:"required,max=255"`
}

func (p *Plugin) createTag(c *axe.Context) {

	var fm fmTag
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	t := Tag{Name: fm.Name}
	if err := p.Db.Create(&t).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, t)

}

func (p *Plugin) showTag(c *axe.Context) {

	var tag Tag
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&tag).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	if err := p.Db.Model(&tag).Association("Articles").Find(&tag.Articles).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, tag)
}

func (p *Plugin) updateTag(c *axe.Context) {

	var tag Tag
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&tag).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	var fm fmTag
	if err := c.Bind(&fm); err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	if err := p.Db.Model(&tag).Update("name", fm.Name).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, tag)

}

func (p *Plugin) destroyTag(c *axe.Context) {
	var tag Tag
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&tag).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	if err := p.Db.Model(&tag).Association("Articles").Clear().Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}

	if err := p.Db.Delete(&tag).Error; err != nil {
		c.Abort(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, tag)

}
