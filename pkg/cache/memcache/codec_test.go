package memcache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	t.Run("UnsupportedType", func(t *testing.T) {
		c := &defaultCodec{}
		orig := struct {
			FirstField  string
			SecondField string
		}{
			FirstField:  "first_field",
			SecondField: "second_field",
		}

		res, err := c.Marshal(orig)

		assert.ErrorIs(t, err, errUnsupportedValue)
		assert.Nil(t, res)
	})

	t.Run("Bytes", func(t *testing.T) {
		c := &defaultCodec{}
		orig := []byte{1, 2, 3}

		res, err := c.Marshal(orig)

		assert.NoError(t, err)
		assert.Equal(t, orig, res)
	})
}

func TestUnmarshal(t *testing.T) {
	t.Run("ReturnSame", func(t *testing.T) {
		c := &defaultCodec{}
		orig := []byte{1, 2, 3}

		p, err := c.Unmarshal(orig)

		assert.NoError(t, err)
		assert.Equal(t, orig, p)
	})
}
