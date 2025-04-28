package main

import (
	"expvar"
	"runtime"
	"time"

	"github.com/Gumilho/gumi-fetch/internal/controller"
	"github.com/Gumilho/gumi-fetch/internal/database"
	"github.com/Gumilho/gumi-fetch/internal/env"
	"github.com/Gumilho/gumi-fetch/internal/store"
	"github.com/Gumilho/gumi-fetch/internal/types"
	"go.uber.org/zap"
)

const version = "0.0.1"

type application struct {
	config         config
	logger         *zap.SugaredLogger
	showController types.Controller
	malController  controller.MALController
}

type config struct {
	addr   string
	apiURL string
	env    string
	db     dbConfig
	mal    malConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  time.Duration
}

type malConfig struct {
	clientID     string
	clientSecret string
}

// @title						GumiFetch
// @description				Backend API for GumiFetch project
// @contact.name				Gumi
// @contact.url				https://gumilho.com
// @contact.email				gumilho2@gmail.com
// @license.name				MIT
// @license.url				https://opensource.org/licenses/MIT
// @BasePath					/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {
	// Config
	cfg := config{
		addr:   env.GetString("ADDR", ":8000"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:8000"),
		env:    env.GetString("ENV", "development"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/auth?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetDuration("DB_MAX_IDLE_TIME", 15*time.Minute),
		},
		mal: malConfig{
			clientID:     env.GetString("MAL_CLIENT_ID", ""),
			clientSecret: env.GetString("MAL_CLIENT_SECRET", ""),
		},
	}

	// Instantiate the dependencies

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	if cfg.env == "development" {
		logger = zap.Must(zap.NewDevelopment()).Sugar()
	}
	defer logger.Sync()
	// Database
	db, err := database.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		logger.Panic(err)
	}
	defer db.Close()
	logger.Info("database successfully connected")

	// Store
	showStore := store.NewShowStore(db)

	// Controller
	showController := controller.NewShowController(showStore, logger)
	malController := controller.NewMALController(logger, cfg.mal.clientID)

	app := application{
		config:         cfg,
		logger:         logger,
		showController: showController,
		malController:  *malController,
	}
	// Metrics collected
	expvar.NewString("version").Set(version)
	expvar.Publish("database", expvar.Func(func() any {
		return db.Stats()
	}))
	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))

	mux := app.mount()
	logger.Fatal(app.run(mux))
}
