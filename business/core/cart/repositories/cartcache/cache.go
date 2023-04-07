package cartcache

import (
	"mergedup/business/core/cart"
	"strconv"
	"sync"

	"go.uber.org/zap"
)

// Store manages the set of APIs for user data and caching.
type Store struct {
	log    *zap.SugaredLogger
	storer cart.Storer
	cache  map[string]*cart.Cart
	mu     sync.RWMutex
}

// NewStore constructs the api for data and caching access.
func NewStore(log *zap.SugaredLogger) *Store {
	return &Store{
		log:    log,
		cache:  map[string]*cart.Cart{},
	}
}
// =============================================================================

// readCache performs a safe search in the cache for the specified key.
func (s *Store) readCache(key string) (cart.Cart, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	crt, exists := s.cache[key]
	if !exists {
		return cart.Cart{}, false
	}

	return *crt, true
}

// writeCache performs a safe write to the cache for the specified user.
func (s *Store) writeCache(crt cart.Cart) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.cache[strconv.Itoa(int(crt.ID))] = &crt
}

// deleteCache performs a safe removal from the cache for the specified user.
func (s *Store) deleteCache(crt cart.Cart) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.cache, strconv.Itoa(int(crt.ID)))
}
