package web

import (
	"fmt"
)

// H hash
type H map[string]interface{}

// K key
type K string

// HTTP http
type HTTP struct {
	Name   string
	Port   int
	Theme  string
	Secure bool
}

// PostgreSQL postgresql
type PostgreSQL struct {
	Host     string
	Port     int
	DbName   string
	User     string
	Password string
	SslMode  string
}

// DataSource datasource url
func (p *PostgreSQL) DataSource() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s?sslmode=%s",
		"postgres",
		p.User,
		p.Password,
		p.Host,
		p.Port,
		p.DbName,
		p.SslMode,
	)
}

// Redis redis
type Redis struct {
	Host string
	Port int
	Db   int
}

// RabbitMQ rabbitmq
type RabbitMQ struct {
	Host     string
	Port     string
	User     string
	Password string
	Virtual  string
}

// Secrets secrets
type Secrets struct {
	Hmac   string
	Aes    string
	Cookie string
	Jwt    string
}

// Configuration configuration
type Configuration struct {
	Env        string
	PostgreSQL PostgreSQL
	Redis      Redis
	RabbitMQ   RabbitMQ
	HTTP       HTTP
	Secrets    Secrets
}

// Home home url
func (p *Configuration) Home() string {
	if p.IsProduction() {
		scheme := "http"
		if p.HTTP.Secure {
			scheme += "s"
		}
		return scheme + "://" + p.HTTP.Name
	}
	return fmt.Sprintf("http://localhost:%d", p.HTTP.Port)
}

// IsProduction production mode ?
func (p *Configuration) IsProduction() bool {
	return p.Env == "production"
}
