package itemgrp

import (
	"mergedup/business/auth"
	"mergedup/business/core/item"
)

// Handlers manages the set of item endpoints.
type Handlers struct {
	Item *item.Core
	Auth *auth.Auth
}
