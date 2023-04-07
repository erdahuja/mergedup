package cartitemdb

import (
	"mergedup/business/core/cartitem"
	"time"
)

// dbItem represents an individual cartitem.Cart
type dbCartItem struct {
	ID          int64     `db:"id"`           // Unique identifier.
	CartID      int64     `db:"cart_id"`      // Display name of the cartitem.Cart
	ItemID      int64     `db:"item_id"`      // Price for one item in cents.
	Quantity    int       `db:"quantity"`     // Original number of items available.
	DateCreated time.Time `db:"date_created"` // When the item was added.
	DateUpdated time.Time `db:"date_updated"` // When the item record was last modified.
}

func toDBItem(prd cartitem.CartItem) dbCartItem {
	prdDB := dbCartItem{
		CartID:      prd.CartID,
		ItemID:      prd.ItemID,
		Quantity:    prd.Quantity,
		DateCreated: prd.DateCreated.UTC(),
		DateUpdated: prd.DateUpdated.UTC(),
	}

	return prdDB
}

func toCoreItem(dbItm dbCartItem) cartitem.CartItem {
	prd := cartitem.CartItem{
		ID:          dbItm.ID,
		CartID:      dbItm.CartID,
		ItemID:      dbItm.ItemID,
		Quantity:    dbItm.Quantity,
		DateCreated: dbItm.DateCreated.In(time.Local),
		DateUpdated: dbItm.DateUpdated.In(time.Local),
	}

	return prd
}

func toCoreItemSlice(dbItems []dbCartItem) []cartitem.CartItem {
	prds := make([]cartitem.CartItem, len(dbItems))
	for i, dbPrd := range dbItems {
		prds[i] = toCoreItem(dbPrd)
	}
	return prds
}
