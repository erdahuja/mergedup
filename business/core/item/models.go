package item

import (
	"time"
)

// Item represents an individual item.
type Item struct {
	ID          int64
	Name        string
	Cost        int
	Quantity    int
	Sold        int
	Revenue     int
	UserID      int64
	DateCreated time.Time
	DateUpdated time.Time
}

// NewProduct is what we require from clients when adding a Item.
type NewItem struct {
	Name     string
	Cost     int
	Quantity int
}

// UpdateItem have fields which clients can send just the fields to update
type UpdateItem struct {
	Quantity *int `json:"quantity" validate:"omitempty"`
}
