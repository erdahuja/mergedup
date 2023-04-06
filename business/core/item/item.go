// Package item contains item related CRUD functionality.
package item

import (
	"context"
	"mergedup/foundation/config"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Item manages the set of API's for Item access.
type Item struct {
	log   *zap.SugaredLogger
	db    *sqlx.DB
	name  string
	table string
}

// New constructs a record for api access.
func New(dbname, table string,
	log *zap.SugaredLogger,
	db *sqlx.DB,
	mainCfg *config.Configurations,
) (Item, error) {

	return Item{
		log:   log,
		db:    db,
		name:  dbname,
		table: table,
	}, nil

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
	Create(ctx context.Context, itm Item) error
	Delete(ctx context.Context, itm Item) error
	Query(ctx context.Context, filter QueryFilter) ([]Item, error)
	QueryByID(ctx context.Context, ItemID uuid.UUID) (Item, error)
}

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
