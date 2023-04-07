package cartdb

import (
	"mergedup/business/core/cart"
	"time"
)

// Cart represents an individual cart row.
type dbCart struct {
	ID          int64
	UserID      int64
	DateCreated time.Time
	DateUpdated time.Time
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
