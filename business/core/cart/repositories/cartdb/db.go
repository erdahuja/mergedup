package cartdb

import (
	"context"
	"fmt"
	"mergedup/business/core/cart"
	"mergedup/business/sys/database"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Store manages the set of APIs for user database access.
type Store struct {
	log    *zap.SugaredLogger
	db     sqlx.ExtContext
	inTran bool
}

// NewStore constructs the api for data access.
func NewStore(log *zap.SugaredLogger, db *sqlx.DB) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}

// WithinTran runs passed function and do commit/rollback at the end.
func (s *Store) WithinTran(ctx context.Context, fn func(s cart.Storer) error) error {
	if s.inTran {
		return fn(s)
	}

	f := func(tx *sqlx.Tx) error {
		s := &Store{
			log:    s.log,
			db:     tx,
			inTran: true,
		}
		return fn(s)
	}

	return database.WithinTran(ctx, s.log, s.db.(*sqlx.DB), f)
}

// Create adds a Item to the database. It returns the created Item with
// fields like ID and DateCreated populated.
func (s *Store) Create(ctx context.Context, cart cart.Cart) error {
	const q = `
	INSERT INTO cart
		(user_id, date_created, date_updated)
	VALUES
		(:user_id, :date_created, :date_updated)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, toDBCart(cart)); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}
