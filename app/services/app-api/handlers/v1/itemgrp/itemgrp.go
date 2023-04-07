package itemgrp

import (
	"context"
	"errors"
	"fmt"
	"mergedup/business/auth"
	"mergedup/business/core/item"
	"mergedup/business/core/user"
	"mergedup/business/sys/validate"
	"mergedup/foundation/web"
	"net/http"
	"strconv"
)

// Handlers manages the set of item endpoints.
type Handlers struct {
	Item *item.Core
	Auth *auth.Auth
}

// Create adds a new user to the system.
func (h Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var nu item.NewItem
	if err := web.Decode(r, &nu); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	usr, err := h.Item.Create(ctx, nu)
	if err != nil {
		if errors.Is(err, user.ErrUniqueEmail) {
			return web.NewRequestError(err, http.StatusConflict)
		}
		return fmt.Errorf("user[%+v]: %w", &usr, err)
	}

	return web.Respond(ctx, w, usr, http.StatusCreated)
}

// Query returns a list of users with paging.
func (h Handlers) Query(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	filter, err := parseFilter(r)
	if err != nil {
		return err
	}
	users, err := h.Item.Query(ctx, filter)
	if err != nil {
		return fmt.Errorf("unable to query for users: %w", err)
	}
	return web.Respond(ctx, w, users, http.StatusOK)
}

func parseFilter(r *http.Request) (item.QueryFilter, error) {
	values := r.URL.Query()

	var filter item.QueryFilter

	if quantity := values.Get("quantity"); quantity != "" {
		qua, err := strconv.ParseInt(quantity, 10, 64)
		if err != nil {
			return item.QueryFilter{}, validate.NewFieldsError("quantity", err)
		}
		filter.WithQuantity(int(qua))
	}

	return filter, nil
}
