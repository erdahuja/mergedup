// cart api system
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
	"mergedup/foundation/validate"
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

	crt, err := h.Cart.QueryByID(ctx, nu.CartID)
	if err != nil {
		return web.NewRequestError(errors.New("unable to find itm:"+err.Error()), http.StatusNotFound)
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

	crti, err := h.CartItem.Create(ctx, nu)
	if err != nil {
		// TODO - revert quantity of items
		return fmt.Errorf("user[%+v]: %w", &crt, err)
	}

	return web.Respond(ctx, w, crti, http.StatusCreated)
}

// QueryByID returns a user by its ID.
func (h Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	cartID := web.Param(r, "cartID")
	if cartID == "" {
		return web.NewRequestError(ErrInvalidID, http.StatusBadRequest)
	}

	id, err := strconv.Atoi(cartID)
	if err != nil {
		return web.NewRequestError(ErrInvalidID, http.StatusBadRequest)
	}

	crt, err := h.Cart.QueryByID(ctx, int64(id))
	if err != nil {
		return web.NewRequestError(ErrInvalidID, http.StatusBadRequest)
	}
	claims := auth.GetClaims(ctx)
	if strconv.Itoa(int(crt.UserID)) != claims.Subject {
		return web.NewRequestError(auth.ErrForbidden, http.StatusForbidden)
	}
	crtItems, err := h.CartItem.QueryByCartID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrNotFound):
			return web.NewRequestError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("ID[%s]: %w", cartID, err)
		}
	}

	return web.Respond(ctx, w, crtItems, http.StatusOK)
}

// CreateCartItem deleted a new item to the cart.
func (h Handlers) DeleteItem(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var nu cartitem.DeleteCartItem
	if err := web.Decode(r, &nu); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	itm, err := h.Item.QueryByID(ctx, nu.ItemID)
	if err != nil {
		return web.NewRequestError(errors.New("unable to find itm:"+err.Error()), http.StatusNotFound)
	}

	c_itms, err := h.CartItem.QueryByCartID(ctx, int(nu.CartID))
	if err != nil {
		// TODO - revert quantity of items
		return fmt.Errorf("user[%+v]: %w", &itm, err)
	}

	var targetCartItem cartitem.CartItem
	for _, itm := range c_itms {
		if itm.ItemID == nu.ItemID {
			targetCartItem = itm
			break
		}
	}

	if targetCartItem.ID == 0 {
		return validate.NewRequestError(ErrInvalidID, http.StatusBadRequest)
	}

	err = h.CartItem.Delete(ctx, targetCartItem)
	if err != nil {
		return fmt.Errorf("user[%+v]: %w", &itm, err)
	}

	oq := itm.Quantity
	uq := targetCartItem.Quantity + oq
	if _, err := h.Item.Update(ctx, itm, item.UpdateItem{
		Quantity: &uq,
	}); err != nil {
		return fmt.Errorf("unable to update itm: %w", err)
	}

	return web.Respond(ctx, w, itm, http.StatusCreated)
}
