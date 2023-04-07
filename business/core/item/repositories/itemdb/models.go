package itemdb

import (
	"time"

	"mergedup/business/core/item"
)

// dbItem represents an individual item.
type dbItem struct {
	ID          int64     `db:"id"`           // Unique identifier.
	Name        string    `db:"name"`         // Display name of the item.
	Cost        int       `db:"cost"`         // Price for one item in cents.
	Quantity    int       `db:"quantity"`     // Original number of items available.
	DateCreated time.Time `db:"date_created"` // When the item was added.
	DateUpdated time.Time `db:"date_updated"` // When the item record was last modified.
}

func toDBItem(prd item.Item) dbItem {
	prdDB := dbItem{
		Name:        prd.Name,
		Cost:        prd.Cost,
		Quantity:    prd.Quantity,
		DateCreated: prd.DateCreated.UTC(),
		DateUpdated: prd.DateUpdated.UTC(),
	}

	return prdDB
}

func toCoreItem(dbItm dbItem) item.Item {
	prd := item.Item{
		ID:          dbItm.ID,
		Name:        dbItm.Name,
		Cost:        dbItm.Cost,
		Quantity:    dbItm.Quantity,
		DateCreated: dbItm.DateCreated.In(time.Local),
		DateUpdated: dbItm.DateUpdated.In(time.Local),
	}

	return prd
}

func toCoreItemSlice(dbItems []dbItem) []item.Item {
	prds := make([]item.Item, len(dbItems))
	for i, dbPrd := range dbItems {
		prds[i] = toCoreItem(dbPrd)
	}
	return prds
}
