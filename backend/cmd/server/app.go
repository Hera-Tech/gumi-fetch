package main

import (
	"expvar"
	"fmt"
	"net/http"

	"github.com/Gumilho/gumi-fetch/docs" // This is required to generate swagger docs
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()
	v1 := http.NewServeMux()
	app.showController.RegisterRoutes(v1)
	app.malController.RegisterRoutes(v1)
	v1.HandleFunc("/debug/vars", expvar.Handler().ServeHTTP)

	docsURL := fmt.Sprintf("http://%s/v1/swagger/doc.json", app.config.apiURL)
	v1.Handle("/swagger/", httpSwagger.Handler(httpSwagger.URL(docsURL)))
	mux.Handle("/v1/", http.StripPrefix("/v1", v1))

	return mux
}

func (app *application) run(mux *http.ServeMux) error {
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = app.config.apiURL
	docs.SwaggerInfo.BasePath = "/v1"

	srv := http.Server{
		Addr:    app.config.addr,
		Handler: app.withLogging(mux),
	}
	app.logger.Infof("server running on addr %s", app.config.addr)
	return srv.ListenAndServe()

}
