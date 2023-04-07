package user

import (
	"net/mail"
	"time"
)

// User represents an individual user in db
type User struct {
	ID           int64   `json:"id"`
	Name         string       `json:"name"`
	Email        mail.Address `json:"email"`
	Roles        []Role       `json:"roles"`
	PasswordHash []byte       `json:"-"`
	Active       bool         `json:"active"`
	DateCreated  time.Time    `json:"dateCreated"`
	DateUpdated  time.Time    `json:"dateUpdated"`
}

// NewUser contains information needed to create a new User.
type NewUser struct {
	Name            string       `json:"name" validate:"required"`
	Email           mail.Address `json:"email" validate:"required,email"`
	Roles           []Role       `json:"roles" validate:"required"`
	Password        string       `json:"password" validate:"required"`
	PasswordConfirm string       `json:"passwordConfirm" validate:"eqfield=Password"`
}

// UpdateUser defines what information may be provided to modify an existing
// User. All fields are optional so clients can send just the fields they want
// changed. It uses pointer fields so we can differentiate between a field that
// was not provided and a field that was provided as explicitly blank. Normally
// we do not want to use pointers to basic types but we make exceptions around
// marshalling/unmarshalling.
type UpdateUser struct {
	Name            *string       `json:"name"`
	Email           *mail.Address `json:"email" validate:"omitempty,email"`
	Roles           []Role        `json:"roles"`
	Password        *string       `json:"password"`
	PasswordConfirm *string       `json:"passwordConfirm" validate:"omitempty,eqfield=Password"`
	Active          *bool         `json:"active"`
}
