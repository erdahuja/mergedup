package itemdb

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mergedup/business/core/item"
	"mergedup/business/sys/database"
	"strings"

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

// Create adds a Item to the database.
func (s *Store) Create(ctx context.Context, itm item.Item) error {
	const q = `
	INSERT INTO items
		(name, cost, quantity, date_created, date_updated)
	VALUES
		(:name, :cost, :quantity, :date_created, :date_updated)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, toDBItem(itm)); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Query gets all Products from the database.
func (s *Store) Query(ctx context.Context, filter item.QueryFilter) ([]item.Item, error) {

	const q = `SELECT * FROM items`

	buf := bytes.NewBufferString(q)
	s.applyFilter(filter, buf)

	var dbItms []dbItem
	if err := database.NamedQuerySlice(ctx, s.log, s.db, buf.String(), struct{}{}, &dbItms); err != nil {
		return nil, fmt.Errorf("namedqueryslice: %w", err)
	}

	return toCoreItemSlice(dbItms), nil
}

// Query gets all Products from the database.
func (s *Store) Update(ctx context.Context, itm item.Item) error {
	const q = `
	UPDATE
		items
	SET
		"quantity" = :quantity,
		"date_updated" = :date_updated
	WHERE
		id = :id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, toDBItem(itm)); err != nil {
		return fmt.Errorf("updating userID[%d]: %w", itm.ID, err)
	}

	return nil
}

func (s *Store) applyFilter(filter item.QueryFilter, buf *bytes.Buffer) {
	var wc []string

	if filter.Quantity != nil {
		wc = append(wc, "p.quantity > :quantity")
	}

	if len(wc) > 0 {
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(wc, " AND "))
	}
}

// QueryByID gets the specified user from the database.
func (s *Store) QueryByID(ctx context.Context, itemID int64) (item.Item, error) {
	data := struct {
		ItemID int64 `db:"id"`
	}{
		ItemID: itemID,
	}

	const q = `
	SELECT
		*
	FROM
		items
	WHERE
		id = :id`

	var usr dbItem
	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &usr); err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			return item.Item{}, item.ErrNotFound
		}
		return item.Item{}, fmt.Errorf("selecting userID[%q]: %w", itemID, err)
	}

	return toCoreItem(usr), nil
}

// WithinTran runs passed function and do commit/rollback at the end.
func (s *Store) WithinTran(ctx context.Context, fn func(s item.Storer) error) error {
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
