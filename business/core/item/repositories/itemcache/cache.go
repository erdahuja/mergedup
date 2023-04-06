package itemcache

import (
	"mergedup/business/core/item"
	"sync"

	"go.uber.org/zap"
)

// Store manages the set of APIs for user data and caching.
type Store struct {
	log    *zap.SugaredLogger
	storer item.Storer
	cache  map[string]*item.Item
	mu     sync.RWMutex
}

// NewStore constructs the api for data and caching access.
func NewStore(log *zap.SugaredLogger, storer item.Storer) *Store {
	return &Store{
		log:    log,
		storer: storer,
		cache:  map[string]*item.Item{},
	}
}
