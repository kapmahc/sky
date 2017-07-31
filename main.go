package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/kapmahc/sky/plugins/forum"
	_ "github.com/kapmahc/sky/plugins/ops/mail"
	_ "github.com/kapmahc/sky/plugins/ops/vpn"
	_ "github.com/kapmahc/sky/plugins/reading"
	_ "github.com/kapmahc/sky/plugins/site"
	_ "github.com/kapmahc/sky/plugins/survey"
	"github.com/kapmahc/sky/web"
)

func main() {
	if err := web.Main(); err != nil {
		log.Fatal(err)
	}
}
