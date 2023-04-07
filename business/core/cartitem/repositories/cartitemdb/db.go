package cartitemdb

import (
	"context"
	"errors"
	"fmt"
	"mergedup/business/core/cartitem"
	"mergedup/business/sys/database"
	"strconv"

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

// Create adds a Item to the database. It returns the created Item with
// fields like ID and DateCreated populated.
func (s *Store) Create(ctx context.Context, item cartitem.CartItem) error {
	const q = `
	INSERT INTO cart_items
		(cart_id, item_id, quantity, date_created, date_updated)
	VALUES
		(:cart_id, :item_id, :quantity, :date_created, :date_updated)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, toDBItem(item)); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// QueryByID gets the specified user from the database.
func (s *Store) QueryByCartID(ctx context.Context, cartID int) ([]cartitem.CartItem, error) {
	data := struct {
		CartID int `db:"cart_id"`
	}{
		CartID: cartID,
	}

	const q = `
	SELECT
		*
	FROM
		cart_items
	WHERE
		cart_id = :cart_id`

	var ci []dbCartItem
	if err := database.NamedQuerySlice(ctx, s.log, s.db, q, data, &ci); err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			return []cartitem.CartItem{}, err
		}
		return []cartitem.CartItem{}, fmt.Errorf("selecting cartID[%q]: %w", cartID, err)
	}

	return toCoreItemSlice(ci), nil
}

// Delete removes a user from the database.
func (s *Store) Delete(ctx context.Context, ci cartitem.CartItem) error {
	data := struct {
		CartID string `db:"cart_id"`
		ItemID string `db:"item_id"`
	}{
		CartID: strconv.FormatInt(ci.CartID, 10),
		ItemID: strconv.FormatInt(ci.ItemID, 10),
	}

	const q = `
	DELETE FROM
		cart_items
	WHERE
		cart_id = :cart_id AND item_id = :item_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, data); err != nil {
		return fmt.Errorf("deleting cartID[%d]: %w", ci.ID, err)
	}

	return nil
}

// WithinTran runs passed function and do commit/rollback at the end.
func (s *Store) WithinTran(ctx context.Context, fn func(s cartitem.Storer) error) error {
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
