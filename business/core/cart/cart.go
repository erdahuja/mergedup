package cart

import (
	"context"

	"github.com/google/uuid"
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
	ID    *uuid.UUID    `validate:"omitempty,uuid4"`
}

// ByID sets the ID field of the QueryFilter value.
func (f *QueryFilter) ByID(id uuid.UUID) {
	var zero uuid.UUID
	if id != zero {
		f.ID = &id
	}
}

type Storer interface {
	WithinTran(ctx context.Context, fn func(s Storer) error) error
	Create(ctx context.Context, itm Cart) error
	Delete(ctx context.Context, itm Cart) error
	Query(ctx context.Context, filter QueryFilter) ([]Cart, error)
	QueryByID(ctx context.Context, CartID uuid.UUID) (Cart, error)
}