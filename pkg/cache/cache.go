package cache

import (
	"context"
	"time"

	"google.golang.org/appengine/memcache"
)

// ErrCacheMiss ...
var ErrCacheMiss = memcache.ErrCacheMiss

type Cache interface {
	// Get returns value from cache
	Get(ctx context.Context, key interface{}) (value interface{}, err error)
	// Set value to cache
	Set(ctx context.Context, key, value interface{})
	// SetWithExpiration value to cache with expiration (in seconds)
	SetWithExpiration(ctx context.Context, key, value interface{}, expiration time.Duration)
	// Remove from cache
	Remove(ctx context.Context, key interface{})
}
