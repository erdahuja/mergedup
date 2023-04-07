package cartitem

import "time"

// CartItem represents an individual item in the cart.
type CartItem struct {
	ID          int64
	CartID      int64
	ItemID      int64
	Quantity    int
	DateCreated time.Time
	DateUpdated time.Time
}

// NewCartItem is what we require from clients when adding a Item to cart.
type NewCartItem struct {
	CartID   int64
	ItemID   int64
	Quantity int
}

// DeleteCartItem is what we require from clients when deleting a Item to cart.
type DeleteCartItem struct {
	CartID   int64
	ItemID   int64
}
