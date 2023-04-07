package item

import (
	"time"
)

// Item represents an individual item.
type Item struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Cost        int       `json:"cost"`
	Quantity    int       `json:"quantity"`
	DateCreated time.Time `json:"dateCreated"`
	DateUpdated time.Time `json:"dateUpdated"`
}

// NewProduct is what we require from clients when adding a Item.
type NewItem struct {
	Name     string `json:"name"`
	Cost     int    `json:"cost"`
	Quantity int    `json:"quantity"`
}

// UpdateItem have fields which clients can send just the fields to update
type UpdateItem struct {
	Quantity *int `json:"quantity" validate:"omitempty"`
}
