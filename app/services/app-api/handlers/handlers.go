// Package handlers contains the full set of handler functions and routes
// supported by the web api.
package handlers

import (
	"net/http"
	"os"

	v1 "mergedup/app/services/app-api/handlers/v1"
	"mergedup/foundation/web"

	"mergedup/business/mid"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
	DB       *sqlx.DB
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig) http.Handler {
	var app *web.App

	if app == nil {
		app = web.NewApp(
			cfg.Shutdown,
			mid.Logger(cfg.Log),
			mid.Errors(cfg.Log),
			mid.Metrics(),
			mid.Panics(),
		)
	}

	v1.Register(app, v1.Config{
		Log: cfg.Log,
		DB:  cfg.DB,
	})

	return app
}
