package site

import (
	"os"
	"path"

	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/sky/plugins/auth"
	"github.com/kapmahc/sky/web"
	"github.com/kapmahc/sky/web/cache"
	"github.com/kapmahc/sky/web/i18n"
	"github.com/kapmahc/sky/web/job"
	"github.com/kapmahc/sky/web/settings"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
	"golang.org/x/tools/blog/atom"
)

// Plugin plugin
type Plugin struct {
	Dao      *auth.Dao          `inject:""`
	Db       *gorm.DB           `inject:""`
	Jwt      *auth.Jwt          `inject:""`
	I18n     *i18n.I18n         `inject:""`
	Settings *settings.Settings `inject:""`
	Server   *job.Server        `inject:""`
	Cache    *cache.Cache       `inject:""`
	Render   *render.Render     `inject:""`
	Wrapper  *web.Wrapper       `inject:""`
}

// Init load config
func (p *Plugin) Init() {}

// Atom rss.atom
func (p *Plugin) Atom(lang string) ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Plugin) Sitemap() ([]stm.URL, error) {
	return []stm.URL{}, nil
}

// Workers job handler
func (p *Plugin) Workers() map[string]job.Handler {
	return map[string]job.Handler{}
}

func init() {
	pwd, _ := os.Getwd()
	viper.SetDefault("uploader", map[string]interface{}{
		"dir":  path.Join(pwd, "public", "files"),
		"home": "http://localhost/files",
	})
	viper.SetDefault("redis", map[string]interface{}{
		"host": "localhost",
		"port": 6379,
		"db":   8,
	})

	viper.SetDefault("rabbitmq", map[string]interface{}{
		"user":     "guest",
		"password": "guest",
		"host":     "localhost",
		"port":     "5672",
		"virtual":  "sky-dev",
	})

	viper.SetDefault("database", map[string]interface{}{
		"driver": "postgres",
		"args": map[string]interface{}{
			"host":     "localhost",
			"port":     5432,
			"user":     "postgres",
			"password": "",
			"dbname":   "sky_dev",
			"sslmode":  "disable",
		},
		"pool": map[string]int{
			"max_open": 180,
			"max_idle": 6,
		},
	})

	viper.SetDefault("server", map[string]interface{}{
		"port":  3000,
		"ssl":   false,
		"name":  "www.change-me.com",
		"theme": "bootstrap",
	})

	viper.SetDefault("secrets", map[string]interface{}{
		"jwt":    web.Random(32),
		"aes":    web.Random(32),
		"hmac":   web.Random(32),
		"cookie": web.Random(32),
		"csrf":   web.Random(32),
	})

	viper.SetDefault("elasticsearch", map[string]interface{}{
		"host": "localhost",
		"port": 9200,
	})

	web.Register(&Plugin{})
}
