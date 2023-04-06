package v1

import (
	"net/http"

	"mergedup/app/services/app-api/handlers/v1/cartgrp"
	"mergedup/app/services/app-api/handlers/v1/itemgrp"
	"mergedup/app/services/app-api/handlers/v1/usergrp"
	"mergedup/business/auth"
	"mergedup/business/core/cart"
	"mergedup/business/core/item"
	"mergedup/business/core/user"
	"mergedup/business/mid"
	"mergedup/foundation/web"

	"mergedup/business/core/user/repositories/usercache"
	"mergedup/business/core/user/repositories/userdb"

	"mergedup/business/core/item/repositories/itemcache"
	"mergedup/business/core/item/repositories/itemdb"

	"mergedup/business/core/cart/repositories/cartcache"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log  *zap.SugaredLogger
	Auth *auth.Auth
	DB   *sqlx.DB
}

func Register(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.Auth)
	ruleAdmin := mid.Authorize(cfg.Auth, auth.RuleAdminOnly)
	ruleAny := mid.Authorize(cfg.Auth, auth.RuleAny)

	ugh := usergrp.Handlers{
		User: user.NewCore(usercache.NewStore(cfg.Log, userdb.NewStore(cfg.Log, cfg.DB))),
		Auth: cfg.Auth,
	}
	app.Handle(http.MethodPost, version, "/users", ugh.Create, authen, ruleAdmin)
	app.Handle(http.MethodGet, version, "/users/:id", ugh.QueryByID, authen, ruleAny)
	app.Handle(http.MethodGet, version, "/status", ugh.Status)
	app.Handle(http.MethodGet, version, "/users/token/:kid", ugh.Token)
	app.Handle(http.MethodGet, version, "/users", ugh.Query, authen, ruleAdmin)
	app.Handle(http.MethodPut, version, "/users/:id", ugh.Update, authen, ruleAny)
	app.Handle(http.MethodDelete, version, "/users/:id", ugh.Delete, authen, ruleAny)

	pgh := itemgrp.Handlers{
		Item: item.NewCore(itemcache.NewStore(cfg.Log, itemdb.NewStore(cfg.Log, cfg.DB))),
		Auth:    cfg.Auth,
	}
	app.Handle(http.MethodGet, version, "/items/:page/:rows", pgh.Query, authen)
	app.Handle(http.MethodGet, version, "/items/:id", pgh.QueryByID, authen)
	app.Handle(http.MethodPost, version, "/items", pgh.Create, authen)
	app.Handle(http.MethodPut, version, "/items/:id", pgh.Update, authen)
	app.Handle(http.MethodDelete, version, "/items/:id", pgh.Delete, authen)

	cgh := cartgrp.Handlers{
		Cart: cart.NewCore(cartcache.NewStore(cfg.Log)),
		Auth:    cfg.Auth,
	}
	app.Handle(http.MethodGet, version, "/cart/:page/:rows", cgh.Query, authen)
	app.Handle(http.MethodGet, version, "/cart/:id", cgh.QueryByID, authen)
	app.Handle(http.MethodPost, version, "/cart", cgh.Create, authen)
	app.Handle(http.MethodPut, version, "/cart/:id", cgh.Update, authen)
	app.Handle(http.MethodDelete, version, "/cart/:id", cgh.Delete, authen)
}
