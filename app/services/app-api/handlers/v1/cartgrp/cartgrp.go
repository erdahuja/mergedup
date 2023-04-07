package cartgrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"mergedup/business/auth"
	"mergedup/business/core/cart"
	"mergedup/business/core/cartitem"
	"mergedup/business/core/item"
	"mergedup/business/core/user"
	"mergedup/foundation/web"
)

// Set of error variables for handling user group errors.
var (
	ErrInvalidID = errors.New("ID is not in its proper form")
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	Cart     *cart.Core
	CartItem *cartitem.Core
	Item     *item.Core
	Auth     *auth.Auth
}

// Create adds a new cart to the system.
func (h Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var nu cart.NewCart
	if err := web.Decode(r, &nu); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	crt, err := h.Cart.Create(ctx, nu)
	if err != nil {
		return fmt.Errorf("user[%+v]: %w", &crt, err)
	}

	return web.Respond(ctx, w, crt, http.StatusCreated)
}

// CreateCartItem adds a new item to the cart.
func (h Handlers) CreateCartItem(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var nu cartitem.NewCartItem
	if err := web.Decode(r, &nu); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	itm, err := h.Item.QueryByID(ctx, nu.ItemID)
	if err != nil {
		return web.NewRequestError(errors.New("unable to find itm:"+err.Error()), http.StatusNotFound)
	}
	q := itm.Quantity
	nq := nu.Quantity
	if q < nq {
		return web.NewRequestError(errors.New("insufficient quantity"), http.StatusBadRequest)
	}
	uq := q - nq
	if _, err := h.Item.Update(ctx, itm, item.UpdateItem{
		Quantity: &uq,
	}); err != nil {
		return fmt.Errorf("unable to update itm: %w", err)
	}

	crt, err := h.CartItem.Create(ctx, nu)
	if err != nil {
		// TODO - revert quantity of items
		return fmt.Errorf("user[%+v]: %w", &crt, err)
	}

	return web.Respond(ctx, w, crt, http.StatusCreated)
}

// QueryByID returns a user by its ID.
func (h Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	cartID := web.Param(r, "id")
	if cartID == "" {
		return web.NewRequestError(ErrInvalidID, http.StatusBadRequest)
	}

	id, err := strconv.Atoi(cartID)
	if err != nil {
		return web.NewRequestError(ErrInvalidID, http.StatusBadRequest)
	}

	usr, err := h.CartItem.QueryByCartID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrNotFound):
			return web.NewRequestError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("ID[%s]: %w", cartID, err)
		}
	}

	return web.Respond(ctx, w, usr, http.StatusOK)
}

// CreateCartItem adds a new item to the cart.
func (h Handlers) DeleteItem(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var nu cartitem.DeleteCartItem
	if err := web.Decode(r, &nu); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	itm, err := h.Item.QueryByID(ctx, nu.ItemID)
	if err != nil {
		return web.NewRequestError(errors.New("unable to find itm:"+err.Error()), http.StatusNotFound)
	}

	c_itm, err := h.CartItem.QueryByCartID(ctx, int(nu.CartID))
	if err != nil {
		// TODO - revert quantity of items
		return fmt.Errorf("user[%+v]: %w", &itm, err)
	}

	err = h.CartItem.Delete(ctx, c_itm)
	if err != nil {
		return fmt.Errorf("user[%+v]: %w", &itm, err)
	}

	oq := itm.Quantity
	uq := c_itm.Quantity + oq
	if _, err := h.Item.Update(ctx, itm, item.UpdateItem{
		Quantity: &uq,
	}); err != nil {
		return fmt.Errorf("unable to update itm: %w", err)
	}

	return web.Respond(ctx, w, itm, http.StatusCreated)
}
