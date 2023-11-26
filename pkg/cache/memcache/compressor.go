package memcache

import (
	"bytes"
	"compress/gzip"
	"io"
)

var (
	_ compressor = &dummyCompressor{}
	_ compressor = &gzipCompressor{}
)

type compressor interface {
	// Compress returns compressed bytes
	Compress(payload []byte) ([]byte, error)
	// Decompress returns decompressed bytes
	Decompress(payload []byte) ([]byte, error)
}

type dummyCompressor struct{}

// Compress returns compressed bytes
func (*dummyCompressor) Compress(payload []byte) ([]byte, error) {
	return payload, nil
}

// Decompress returns decompressed bytes
func (*dummyCompressor) Decompress(payload []byte) ([]byte, error) {
	return payload, nil
}

type gzipCompressor struct{}

// Decompress returns decompressed bytes
func (c *gzipCompressor) Decompress(payload []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	return io.ReadAll(r)
}

// Compress returns compressed bytes
func (c *gzipCompressor) Compress(payload []byte) ([]byte, error) {
	b := bytes.NewBuffer(nil)

	w := gzip.NewWriter(b)

	if _, err := w.Write(payload); err != nil {
		return nil, err
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
