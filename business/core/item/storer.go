package item

import "context"

type Storer interface {
	WithinTran(ctx context.Context, fn func(s Storer) error) error
	Create(ctx context.Context, itm Item) error
	Query(ctx context.Context, filter QueryFilter) ([]Item, error)
	QueryByID(ctx context.Context, itemID int64) (Item, error)
	Update(ctx context.Context, itm Item) error
}
