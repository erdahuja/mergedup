package cartitem

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
	Create(ctx context.Context, itm CartItem) error
	Delete(ctx context.Context, itm CartItem) error
	QueryByCartID(ctx context.Context, cartID int) (CartItem, error)
}

// Create adds a Item to the database. It returns the created Item with
// fields like ID and DateCreated populated.
func (c *Core) Create(ctx context.Context, np NewCartItem) (CartItem, error) {
	now := time.Now()
	itm := CartItem{
		CartID:      np.CartID,
		ItemID:      np.ItemID,
		Quantity:    np.Quantity,
		DateCreated: now,
		DateUpdated: now,
	}

	tran := func(s Storer) error {
		if err := s.Create(ctx, itm); err != nil {
			return fmt.Errorf("create: %w", err)
		}
		return nil
	}

	if err := c.storer.WithinTran(ctx, tran); err != nil {
		return CartItem{}, fmt.Errorf("tran: %w", err)
	}

	return itm, nil
}

// QueryByCartID gets the specified cart items from the database.
func (c *Core) QueryByCartID(ctx context.Context, cartID int) (CartItem, error) {
	cart, err := c.storer.QueryByCartID(ctx, cartID)
	if err != nil {
		return CartItem{}, fmt.Errorf("query: %w", err)
	}

	return cart, nil
}

// Delete gets the specified cart items from the database.
func (c *Core) Delete(ctx context.Context, itm CartItem) error {
	err := c.storer.Delete(ctx, itm)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	return nil
}
