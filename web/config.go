package web

import (
	"github.com/spf13/viper"
)

const (
	// LayoutApplication application layout
	LayoutApplication = "layouts/application/index"
)

// Name get server.name
func Name() string {
	return viper.GetString("server.name")
}

// IsProduction production mode ?
func IsProduction() bool {
	return viper.GetString("env") == "production"
}
