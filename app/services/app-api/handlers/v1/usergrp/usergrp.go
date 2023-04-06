// Package usergrp maintains the group of handlers for user access.
package usergrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/mail"
	"time"

	"mergedup/business/auth"
	"mergedup/business/core/user"
	"mergedup/foundation/web"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// Set of error variables for handling user group errors.
var (
	ErrInvalidID = errors.New("ID is not in its proper form")
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	User *user.Core
	Auth *auth.Auth
}

// Create adds a new user to the system.
func (h Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var nu user.NewUser
	if err := web.Decode(r, &nu); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	usr, err := h.User.Create(ctx, nu)
	if err != nil {
		if errors.Is(err, user.ErrUniqueEmail) {
			return web.NewRequestError(err, http.StatusConflict)
		}
		return fmt.Errorf("user[%+v]: %w", &usr, err)
	}

	return web.Respond(ctx, w, usr, http.StatusCreated)
}

// Update updates a user in the system.
func (h Handlers) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var upd user.UpdateUser
	if err := web.Decode(r, &upd); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	userID, err := uuid.Parse(web.Param(r, "id"))
	if err != nil {
		return web.NewRequestError(ErrInvalidID, http.StatusBadRequest)
	}

	claims := auth.GetClaims(ctx)
	if claims.Subject != userID.String() && h.Auth.Authorize(ctx, claims, auth.RuleAdminOnly) != nil {
		return auth.NewAuthError("auth failed")
	}

	usr, err := h.User.QueryByID(ctx, userID)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrNotFound):
			return web.NewRequestError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("ID[%s]: %w", userID, err)
		}
	}

	usr, err = h.User.Update(ctx, usr, upd)
	if err != nil {
		return fmt.Errorf("ID[%s] User[%+v]: %w", userID, &upd, err)
	}

	return web.Respond(ctx, w, usr, http.StatusOK)
}

// Delete removes a user from the system.
func (h Handlers) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID, err := uuid.Parse(web.Param(r, "id"))
	if err != nil {
		return web.NewRequestError(ErrInvalidID, http.StatusBadRequest)
	}

	claims := auth.GetClaims(ctx)
	if claims.Subject != userID.String() && h.Auth.Authorize(ctx, claims, auth.RuleAdminOnly) != nil {
		return auth.NewAuthError("auth failed")
	}

	usr, err := h.User.QueryByID(ctx, userID)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrNotFound):
			return web.Respond(ctx, w, nil, http.StatusNoContent)
		default:
			return fmt.Errorf("ID[%s]: %w", userID, err)
		}
	}

	if err := h.User.Delete(ctx, usr); err != nil {
		return fmt.Errorf("ID[%s]: %w", userID, err)
	}

	return web.Respond(ctx, w, nil, http.StatusNoContent)
}

// Query returns a list of users with paging.
func (h Handlers) Query(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	filter, err := parseFilter(r)
	if err != nil {
		return web.NewRequestError(err, http.StatusBadRequest)
	}

	users, err := h.User.Query(ctx, filter)
	if err != nil {
		return fmt.Errorf("unable to query for users: %w", err)
	}

	return web.Respond(ctx, w, users, http.StatusOK)
}

// QueryByID returns a user by its ID.
func (h Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID, err := uuid.Parse(web.Param(r, "id"))
	if err != nil {
		return web.NewRequestError(ErrInvalidID, http.StatusBadRequest)
	}

	claims := auth.GetClaims(ctx)
	if claims.Subject != userID.String() && h.Auth.Authorize(ctx, claims, auth.RuleAdminOnly) != nil {
		return auth.NewAuthError("auth failed")
	}

	usr, err := h.User.QueryByID(ctx, userID)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrNotFound):
			return web.NewRequestError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("ID[%s]: %w", userID, err)
		}
	}

	return web.Respond(ctx, w, usr, http.StatusOK)
}

// QueryByID returns a user by its ID.
func (h Handlers) Status(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return web.Respond(ctx, w, nil, http.StatusOK)
}

// Token provides an API token for the authenticated user.
func (h Handlers) Token(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	kid := web.Param(r, "kid")
	if kid == "" {
		return web.NewRequestError(errors.New("missing kid"), http.StatusBadRequest)
	}

	email, pass, ok := r.BasicAuth()
	if !ok {
		return auth.NewAuthError("must provide email and password in Basic auth")
	}

	addr, err := mail.ParseAddress(email)
	if err != nil {
		return auth.NewAuthError("invalid email format")
	}

	usr, err := h.User.Authenticate(ctx, *addr, pass)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrNotFound):
			return web.NewRequestError(err, http.StatusNotFound)
		case errors.Is(err, user.ErrAuthenticationFailure):
			return auth.NewAuthError(err.Error())
		default:
			return fmt.Errorf("authenticating: %w", err)
		}
	}

	claims := auth.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   usr.ID.String(),
			Issuer:    "mergedup",
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		},
		Roles: usr.Roles,
	}

	var tkn struct {
		Token string `json:"token"`
	}
	tkn.Token, err = h.Auth.GenerateToken(kid, claims)
	if err != nil {
		return fmt.Errorf("generating token: %w", err)
	}

	return web.Respond(ctx, w, tkn, http.StatusOK)
}

func parseFilter(r *http.Request) (user.QueryFilter, error) {
	values := r.URL.Query()

	var filter user.QueryFilter

	if id, err := uuid.Parse(values.Get("id")); err == nil {
		filter.ByID(id)
	}

	if err := filter.ByName(values.Get("name")); err != nil {
		return user.QueryFilter{}, err
	}

	if email, err := mail.ParseAddress(values.Get("email")); err == nil {
		filter.ByEmail(*email)
	}

	return filter, nil
}