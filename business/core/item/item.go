// Package item contains item related CRUD functionality.
package item

import (
	"context"
	"fmt"
	"mergedup/business/sys/validate"
	"time"
)

// Core manages the set of APIs for item access.
type Core struct {
	storer Storer
}

// NewCore constructs a core for product api access.
func NewCore(storer Storer) *Core {
	return &Core{
		storer: storer,
	}
}

// QueryFilter holds the available fields a query can be filtered on.
type QueryFilter struct {
	Quantity *int `validate:"omitempty,uuid4"`
}

// ByID sets the ID field of the QueryFilter value.
func (f *QueryFilter) ByQuantity(q int) {
	var zero int
	if q != zero {
		f.Quantity = &q
	}
}

// WithQuantity sets the Quantity field of the QueryFilter value.
func (qf *QueryFilter) WithQuantity(quantity int) {
	qf.Quantity = &quantity
}

type Storer interface {
	WithinTran(ctx context.Context, fn func(s Storer) error) error
	Create(ctx context.Context, itm Item) error
	Query(ctx context.Context, filter QueryFilter) ([]Item, error)
	QueryByID(ctx context.Context, itemID int64) (Item, error)
	Update(ctx context.Context, itm Item) error
}

// Create adds a Item to the database. It returns the created Item with
// fields like ID and DateCreated populated.
func (c *Core) Create(ctx context.Context, np NewItem) (Item, error) {
	now := time.Now()
	itm := Item{
		Name:        np.Name,
		Cost:        np.Cost,
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
		return Item{}, fmt.Errorf("tran: %w", err)
	}

	return itm, nil
}

// Query gets all Products from the database.
func (c *Core) Query(ctx context.Context, filter QueryFilter) ([]Item, error) {
	prds, err := c.storer.Query(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return prds, nil
}

// QueryByID gets the specified user from the database.
func (c *Core) QueryByID(ctx context.Context, itemID int64) (Item, error) {
	item, err := c.storer.QueryByID(ctx, itemID)
	if err != nil {
		return Item{}, fmt.Errorf("query: %w", err)
	}

	return item, nil
}

// Update replaces a user document in the database.
func (c *Core) Update(ctx context.Context, itm Item, uu UpdateItem) (Item, error) {
	if err := validate.Check(uu); err != nil {
		return Item{}, fmt.Errorf("validating data: %w", err)
	}

	if uu.Quantity != nil {
		itm.Quantity = *uu.Quantity
	}
	itm.DateUpdated = time.Now()

	tran := func(s Storer) error {
		if err := s.Update(ctx, itm); err != nil {
			return fmt.Errorf("create: %w", err)
		}
		return nil
	}

	if err := c.storer.WithinTran(ctx, tran); err != nil {
		return Item{}, fmt.Errorf("tran: %w", err)
	}

	return itm, nil
}
