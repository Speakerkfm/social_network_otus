package memcache

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKey2String(t *testing.T) {
	for _, tc := range []struct {
		in       interface{}
		expected string
	}{
		{in: "string", expected: "string"},
		{in: float64(100), expected: "100"},
		{in: float32(100), expected: "100"},
		{in: int(100), expected: "100"},
		{in: int64(100), expected: "100"},
		{in: int32(100), expected: "100"},
		{in: int16(100), expected: "100"},
		{in: int8(100), expected: "100"},
		{in: uint(100), expected: "100"},
		{in: uint64(100), expected: "100"},
		{in: uint32(100), expected: "100"},
		{in: uint16(100), expected: "100"},
		{in: uint8(100), expected: "100"},
		{in: []byte("100"), expected: "100"},
		{in: errors.New("fmt.Sprintf"), expected: "fmt.Sprintf"},
	} {
		assert.Equal(t, tc.expected, key2string(tc.in))
	}
}

func TestKeys2Strings(t *testing.T) {
	keys := []interface{}{"test", float64(3.14), 200}
	strings := keys2strings(keys)
	assert.EqualValues(t, []string{"test", "3.14", "200"}, strings)
}
