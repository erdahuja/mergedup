package cart

import (
	"time"
)

// Cart represents an individual user in db
type Cart struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

// NewUser contains information needed to create a new User.
type NewCart struct {
	UserID int64 `json:"user_id" validate:"required"`
}
