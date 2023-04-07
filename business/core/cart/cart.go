package cart

import (
	"context"
	"fmt"
	"time"
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
	Create(ctx context.Context, itm Cart) error
}

func (c *Core) Create(ctx context.Context, np NewCart) (Cart, error) {
	now := time.Now()
	crt := Cart{
		UserID:      np.UserID,
		DateCreated: now,
		DateUpdated: now,
	}

	tran := func(s Storer) error {
		if err := s.Create(ctx, crt); err != nil {
			return fmt.Errorf("create: %w", err)
		}
		return nil
	}

	if err := c.storer.WithinTran(ctx, tran); err != nil {
		return Cart{}, fmt.Errorf("tran: %w", err)
	}

	return crt, nil
}
