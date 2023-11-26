package memcache

import (
	"context"
	"errors"
	"io"
	"strings"
	"time"

	"google.golang.org/appengine/memcache"

	"github.com/Speakerkfm/social_network_otus/pkg/cache"
)

// Options represents options for memcachedCache
type Options struct {
	codec      Codec
	compressor compressor
	ttl        time.Duration
}

// Option represent one option for memcachedCache
type Option func(o *Options)

// WithTTL sets ttl to options
func WithTTL(d time.Duration) Option {
	return func(o *Options) {
		o.ttl = d
	}
}

// WithGzipCompressor sets gzip compressor to options
func WithGzipCompressor() Option {
	return func(o *Options) {
		o.compressor = &gzipCompressor{}
	}
}

// WithCodec sets codec to options
func WithCodec(c Codec) Option {
	return func(o *Options) {
		o.codec = c
	}
}

type client interface {
	Get(ctx context.Context, key string) (*memcache.Item, error)
	Set(ctx context.Context, item *memcache.Item) error
	Delete(ctx context.Context, key string) error
}

// MemcachedCache ...
type MemcachedCache struct {
	client     client
	codec      Codec
	compressor compressor
	entityName string
	ttl        time.Duration
}

// New returns new memcached memcachedCache
func New(entityName string, client client, opts ...Option) *MemcachedCache {
	// create options with default values
	o := Options{
		codec:      &defaultCodec{},
		compressor: &dummyCompressor{},
	}

	for _, opt := range opts {
		opt(&o)
	}

	return &MemcachedCache{
		client:     client,
		entityName: entityName,
		// options
		ttl:        o.ttl,
		codec:      o.codec,
		compressor: o.compressor,
	}
}

func (c *MemcachedCache) get(ctx context.Context, key string) ([]byte, error) {
	item, err := c.client.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	return item.Value, nil
}

func (c *MemcachedCache) set(ctx context.Context, key string, value []byte, expiration time.Duration) error {
	// use default ttl if expiration is not specified for current operation
	if expiration == -1 {
		expiration = c.ttl
	}

	err := c.client.Set(ctx, &memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: expiration,
	})
	if err != nil {
		return err
	}
	return nil
}

// Get returns value from cache
func (c *MemcachedCache) Get(ctx context.Context, key interface{}) (interface{}, error) {

	payload, err := c.get(ctx, key2string(key))
	if err != nil && err != io.EOF { //nolint:errorlint // comparing with io.EOF is fine
		if errors.Is(err, memcache.ErrCacheMiss) {
			return nil, cache.ErrCacheMiss
		}
		return nil, err
	}

	value, err := c.unmarshalFromPayload(key, payload)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// unmarshalFromPayload decompress and unmarshal bytes payload to value
func (c *MemcachedCache) unmarshalFromPayload(key interface{}, payload []byte) (interface{}, error) {

	decompressedPayload, err := c.compressor.Decompress(payload)
	if err != nil {
		return nil, err
	}

	res, err := c.codec.Unmarshal(decompressedPayload)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Set value to cache
func (c *MemcachedCache) Set(ctx context.Context, key, value interface{}) {
	c.SetWithExpiration(ctx, key, value, -1)
}

// SetWithExpiration set value to cache with expiration
func (c *MemcachedCache) SetWithExpiration(ctx context.Context, key, value interface{}, expiration time.Duration) {

	payload, err := c.marshalToPayload(key, value)
	if err != nil {
		return
	}

	err = c.set(ctx, key2string(key), payload, expiration)
	if err != nil {
		return
	}
}

// marshalToPayload marshal and compress value to bytes payload
func (c *MemcachedCache) marshalToPayload(key interface{}, value interface{}) ([]byte, error) {

	payload, err := c.codec.Marshal(value)
	if err != nil {
		return nil, err
	}

	payload, err = c.compressor.Compress(payload)
	if err != nil {
		return nil, err
	}

	return payload, err
}

// Remove from cache
func (c *MemcachedCache) Remove(ctx context.Context, key interface{}) {
	err := c.client.Delete(ctx, key2string(key))
	if err != nil {
		return
	}
}

func isTimeoutError(err error) bool {
	return strings.HasSuffix(err.Error(), "i/o timeout")
}
