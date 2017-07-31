package web

import (
	"fmt"

	"github.com/spf13/viper"
)

// H hash
type H map[string]interface{}

// K key
type K string

// Name get server.name
func Name() string {
	return viper.GetString("server.name")
}

// Home home url
func Home() string {
	if IsProduction() {
		scheme := "http"
		if viper.GetBool("server.ssl") {
			scheme = "https"
		}
		return scheme + "://" + Name()
	}
	return fmt.Sprintf("http://localhost:%d", viper.GetInt("server.port"))
}

// IsProduction production mode ?
func IsProduction() bool {
	return viper.GetString("env") == "production"
}
