package memcache

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var (
	_ compressor = &dummyCompressor{}
	_ compressor = &gzipCompressor{}
)

func genRandomData(size int) []byte {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, size)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return b
}

func TestDummyCompressor_Roundtrip(t *testing.T) {
	origPayload := genRandomData(1024)

	c := &dummyCompressor{}

	p, err := c.Compress(origPayload)
	require.NoError(t, err)
	require.NotNil(t, p)

	payload, err := c.Decompress(p)
	require.NoError(t, err)
	require.NotNil(t, payload)

	assert.Equal(t, origPayload, payload)
}

func TestGzipCompressor_Roundtrip(t *testing.T) {
	origPayload := genRandomData(1024)

	c := &gzipCompressor{}

	p, err := c.Compress(origPayload)
	require.NoError(t, err)
	require.NotNil(t, p)

	payload, err := c.Decompress(p)
	require.NoError(t, err)
	require.NotNil(t, payload)

	assert.Equal(t, origPayload, payload)
}
