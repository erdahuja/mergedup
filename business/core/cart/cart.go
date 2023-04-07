package cart

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrNotFound              = errors.New("cart not found")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

// Core manages the set of APIs for user access.
type Core struct {
	storer Storer
}

// NewCore constructs a core for user api access.
func NewCore(storer Storer) *Core {
	return &Core{
		storer: storer,
	}
}

type Storer interface {
	WithinTran(ctx context.Context, fn func(s Storer) error) error
	Create(ctx context.Context, itm Cart) (Cart, error)
	QueryByID(ctx context.Context, cartID int64) (Cart, error)
}

func (c *Core) Create(ctx context.Context, np NewCart) (Cart, error) {
	now := time.Now()
	crt := Cart{
		UserID:      np.UserID,
		DateCreated: now,
		DateUpdated: now,
	}

	tran := func(s Storer) error {
		crtDB, err := s.Create(ctx, crt)
		if err != nil {
			return fmt.Errorf("create: %w", err)
		}
		crt = crtDB
		return nil
	}

	if err := c.storer.WithinTran(ctx, tran); err != nil {
		return Cart{}, fmt.Errorf("tran: %w", err)
	}

	return crt, nil
}

func (c *Core) QueryByID(ctx context.Context, cartID int64) (Cart, error) {
	crt, err := c.storer.QueryByID(ctx, cartID)
	if err != nil {
		return Cart{}, fmt.Errorf("query: %w", err)
	}

	return crt, nil
}
