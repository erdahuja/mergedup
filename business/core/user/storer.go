package user

import (
	"context"
	"net/mail"
)

//
// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	WithinTran(ctx context.Context, fn func(s Storer) error) error
	Create(ctx context.Context, usr User) error
	Update(ctx context.Context, usr User) error
	Delete(ctx context.Context, usr User) error
	Query(ctx context.Context) ([]User, error) // ideally it should have pagination
	QueryByID(ctx context.Context, userID int) (User, error)
	QueryByEmail(ctx context.Context, email mail.Address) (User, error)
}
