package v1

import (
	"net/http"

	"mergedup/app/services/app-api/handlers/v1/cartgrp"
	"mergedup/app/services/app-api/handlers/v1/itemgrp"
	"mergedup/app/services/app-api/handlers/v1/usergrp"
	"mergedup/business/auth"
	"mergedup/business/core/cart"
	"mergedup/business/core/cart/repositories/cartdb"
	"mergedup/business/core/cartitem"
	"mergedup/business/core/cartitem/repositories/cartitemdb"
	"mergedup/business/core/item"
	"mergedup/business/core/item/repositories/itemdb"
	"mergedup/business/core/user"
	"mergedup/business/mid"
	"mergedup/foundation/web"

	"mergedup/business/core/user/repositories/usercache"
	"mergedup/business/core/user/repositories/userdb"

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
	ruleUser := mid.Authorize(cfg.Auth, auth.RuleUserOnly)

	ugh := usergrp.Handlers{
		User: user.NewCore(usercache.NewStore(cfg.Log, userdb.NewStore(cfg.Log, cfg.DB))),
		Auth: cfg.Auth,
	}
	app.Handle(http.MethodGet, version, "/users/token/:id", ugh.Token)               // public api on basic auth, to get auth token
	app.Handle(http.MethodPost, version, "/users", ugh.Create, authen, ruleAdmin)    // add users (only by admin)
	app.Handle(http.MethodGet, version, "/users", ugh.Query, authen, ruleAdmin)      // all users
	app.Handle(http.MethodPut, version, "/users/:id", ugh.Update, authen, ruleAdmin) // update user
	app.Handle(http.MethodGet, version, "/users/:id", ugh.QueryByID, authen, ruleAny)

	igh := itemgrp.Handlers{
		Item: item.NewCore(itemdb.NewStore(cfg.Log, cfg.DB)),
		Auth: cfg.Auth,
	}
	app.Handle(http.MethodPost, version, "/items", igh.Create, authen, ruleAdmin)
	app.Handle(http.MethodGet, version, "/items", igh.Query, authen, ruleAny)

	cgh := cartgrp.Handlers{
		Cart:     cart.NewCore(cartdb.NewStore(cfg.Log, cfg.DB)),
		CartItem: cartitem.NewCore(cartitemdb.NewStore(cfg.Log, cfg.DB)),
		Item:     item.NewCore(itemdb.NewStore(cfg.Log, cfg.DB)),
		Auth:     cfg.Auth,
	}

	app.Handle(http.MethodPost, version, "/cart", cgh.Create, authen, ruleUser) // add a new cart

	app.Handle(http.MethodPost, version, "/cart-items/:id", cgh.CreateCartItem, authen, ruleUser) // add item in cart
	app.Handle(http.MethodGet, version, "/cart-items/:id", cgh.QueryByID, authen, ruleUser)       // view cart
	app.Handle(http.MethodDelete, version, "/cart-items/:id", cgh.DeleteItem, authen, ruleUser)   // remove itm from cart

	app.Handle(http.MethodGet, version, "/status", ugh.Status)

}
