package cart

import (
	"context"
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

// QueryFilter holds the available fields a query can be filtered on.
type QueryFilter struct {
	ID    *int64   `validate:"omitempty,uuid4"`
}

// ByID sets the ID field of the QueryFilter value.
func (f *QueryFilter) ByID(id int64) {
	var zero int64
	if id != zero {
		f.ID = &id
	}
}

type Storer interface {
	WithinTran(ctx context.Context, fn func(s Storer) error) error
	Create(ctx context.Context, itm Cart) error
	Delete(ctx context.Context, itm Cart) error
	Query(ctx context.Context, filter QueryFilter) ([]Cart, error)
	QueryByID(ctx context.Context, CartID int) (Cart, error)
}