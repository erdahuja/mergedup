package cartdb

import (
	"mergedup/business/core/cart"
	"time"
)

// Cart represents an individual cart row.
type dbCart struct {
	ID          int64     `db:"id"`
	UserID      int64     `db:"user_id"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

// NewCart is what we require from clients when adding a Item.
type NewCart struct {
	UserID int64
}

func toDBCart(cart cart.Cart) dbCart {
	prdDB := dbCart{
		UserID:      cart.UserID,
		DateCreated: cart.DateCreated.UTC(),
		DateUpdated: cart.DateUpdated.UTC(),
	}

	return prdDB
}

func toCoreCart(db dbCart) cart.Cart {
	prdDB := cart.Cart{
		ID:          db.ID,
		UserID:      db.UserID,
		DateCreated: db.DateCreated.UTC(),
		DateUpdated: db.DateUpdated.UTC(),
	}

	return prdDB
}
