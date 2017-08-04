package forum

import (
	"github.com/kapmahc/axe"
)

func (p *Plugin) indexTags(c *axe.Context) (interface{}, error) {
	var tags []Tag
	if err := p.Db.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

type fmTag struct {
	Name string `json:"name" binding:"required,max=255"`
}

func (p *Plugin) createTag(c *axe.Context) (interface{}, error) {

	var fm fmTag
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	t := Tag{Name: fm.Name}
	if err := p.Db.Create(&t).Error; err != nil {
		return nil, err
	}

	return t, nil

}

func (p *Plugin) showTag(c *axe.Context) (interface{}, error) {

	var tag Tag
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&tag).Error; err != nil {
		return nil, err
	}
	if err := p.Db.Model(&tag).Association("Articles").Find(&tag.Articles).Error; err != nil {
		return nil, err
	}
	return tag, nil
}

func (p *Plugin) updateTag(c *axe.Context) (interface{}, error) {

	var tag Tag
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&tag).Error; err != nil {
		return nil, err
	}

	var fm fmTag
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	if err := p.Db.Model(&tag).Update("name", fm.Name).Error; err != nil {
		return nil, err
	}

	return tag, nil

}

func (p *Plugin) destroyTag(c *axe.Context) (interface{}, error) {
	var tag Tag
	if err := p.Db.Where("id = ?", c.Params["id"]).First(&tag).Error; err != nil {
		return nil, err
	}

	if err := p.Db.Model(&tag).Association("Articles").Clear().Error; err != nil {
		return nil, err
	}

	if err := p.Db.Delete(&tag).Error; err != nil {
		return nil, err
	}
	return tag, nil

}
