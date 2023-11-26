package adapterristretto

import (
	"context"
	"log"
	"time"

	"github.com/Speakerkfm/social_network_otus/pkg/cache"
)

type ristretto interface {
	Get(key interface{}) (value interface{}, ok bool)
	Set(key, value interface{}, cost int64) bool
	Del(key interface{})
	SetWithTTL(key, value interface{}, cost int64, ttl time.Duration) bool
}

type adapter struct {
	l ristretto
}

// New returns lru with interface utils/cache
func New(l ristretto) cache.Cache {
	return &adapter{
		l: l,
	}
}

// Get returns value from cache
func (a *adapter) Get(ctx context.Context, key interface{}) (interface{}, error) {
	if isKeyTypeUnsupported(key) {
		return nil, nil
	}
	if v, ok := a.l.Get(key); ok {
		return v, nil
	}
	return nil, nil
}

// GetMulti returns multiple values from cache
func (a *adapter) GetMulti(ctx context.Context, keys []interface{}) (map[interface{}]interface{}, error) {
	items := make(map[interface{}]interface{})
	for _, key := range keys {
		if isKeyTypeUnsupported(key) {
			continue
		}
		if value, ok := a.l.Get(key); ok {
			items[key] = value
		}
	}
	return items, nil
}

// Set value to cache
func (a *adapter) Set(ctx context.Context, key, value interface{}) {
	log.Printf("%+v", value)
	if isKeyTypeUnsupported(key) {
		return
	}
	a.l.Set(key, value, 1)
}

// SetWithExpiration value to cache with expiration
func (a *adapter) SetWithExpiration(ctx context.Context, key, value interface{}, expiration time.Duration) {
	if isKeyTypeUnsupported(key) {
		return
	}
	a.l.SetWithTTL(key, value, 1, expiration)
}

// Remove from cache
func (a *adapter) Remove(ctx context.Context, key interface{}) {
	if isKeyTypeUnsupported(key) {
		return
	}
	a.l.Del(key)
}

func isKeyTypeUnsupported(key interface{}) bool {
	switch key.(type) {
	case uint64, string, []byte, byte, int, int32, uint32, int64:
		return false
	}
	return true
}
