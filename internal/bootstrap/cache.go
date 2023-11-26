package bootstrap

import (
	"github.com/dgraph-io/ristretto"

	"github.com/Speakerkfm/social_network_otus/pkg/cache"
	"github.com/Speakerkfm/social_network_otus/pkg/cache/adapterristretto"
)

func NewCache(entity string) (cache.Cache, error) {
	ristrettoCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		return nil, err
	}
	return adapterristretto.New(ristrettoCache), nil
}
