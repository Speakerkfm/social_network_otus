package memcache

import (
	"errors"
	"fmt"
)

var (
	_ Codec = &defaultCodec{}

	errUnsupportedValue = errors.New("unsupported value type, need []byte")
)

// Codec interface to code/decode structs to memcached
type Codec interface {
	// Unmarshal returns decoded object
	Unmarshal(payload []byte) (interface{}, error)
	// Marshal returns coded object
	Marshal(src interface{}) ([]byte, error)
}

type defaultCodec struct{}

func (c *defaultCodec) Unmarshal(payload []byte) (interface{}, error) {
	return payload, nil
}

func (c *defaultCodec) Marshal(src interface{}) ([]byte, error) {
	if value, ok := src.([]byte); ok {
		return value, nil
	}
	return nil, fmt.Errorf("%T: %w", src, errUnsupportedValue)
}
