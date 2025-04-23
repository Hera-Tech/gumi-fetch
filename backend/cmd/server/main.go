package main

import (
	"time"

	"github.com/Gumilho/gumi-fetch/internal/env"
	"go.uber.org/zap"
)

type application struct {
	config        config
	store         store.Storage
	authenticator auth.Authenticator
	hasher        hashing.Hasher
	logger        *zap.SugaredLogger
	rateLimiter   ratelimiter.Limiter
	cacheStorage  cache.Storage
}

type config struct {
	addr   string
	apiURL string
	env    string
	db     dbConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  time.Duration
}

// @title						GumiFetch
// @description			Backend API for GumiFetch project
// @contact.name		Gumi
// @contact.url			https://gumilho.com
// @contact.email		gumilho2@gmail.com
// @license.name		MIT
// @license.url			https://opensource.org/licenses/MIT
// @BasePath				/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {
	// Config
	cfg := config{
		addr:   env.GetString("ADDR", ":8080"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:8080"),
		env:    env.GetString("ENV", "development"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/auth?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetDuration("DB_MAX_IDLE_TIME", 15*time.Minute),
		},
	}

}
