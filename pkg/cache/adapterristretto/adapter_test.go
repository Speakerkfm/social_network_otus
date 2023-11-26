package adapterristretto

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type stub struct {
	// OnGet callback for Get method
	OnGet func(key interface{}) (value interface{}, ok bool)
	// OnAdd callback for Add method
	OnAdd func(key, value interface{})
	// OnRemove callback for Remove method
	OnRemove func(key interface{}) (present bool)
}

func (s *stub) Get(key interface{}) (value interface{}, ok bool) {
	return s.OnGet(key)
}

func (s *stub) Put(key interface{}, value interface{}) {
	s.OnAdd(key, value)
}

func (s *stub) PutWithTTL(key, value interface{}, ttl time.Duration) {
	s.OnAdd(key, value)
}

func (s *stub) Remove(key interface{}) {
	s.OnRemove(key)
}

func TestNew(t *testing.T) {
	assert.NotNil(t, New(&stub{}))
}

func TestGet(t *testing.T) {
	t.Run("WasCalled", func(t *testing.T) {
		wasCalled := false
		adapter := New(&stub{
			OnGet: func(key interface{}) (value interface{}, ok bool) {
				wasCalled = true
				return nil, true
			},
		})

		v, err := adapter.Get(context.Background(), "key")

		assert.True(t, wasCalled)
		assert.NoError(t, err)
		assert.Nil(t, v)
	})

	t.Run("ReturnsSameResult", func(t *testing.T) {
		adapter := New(&stub{
			OnGet: func(key interface{}) (value interface{}, ok bool) {
				return 1, true
			},
		})

		v, err := adapter.Get(context.Background(), "key")

		assert.NoError(t, err)
		assert.Equal(t, 1, v)
	})

	t.Run("SameParams", func(t *testing.T) {
		adapter := New(&stub{
			OnGet: func(key interface{}) (value interface{}, ok bool) {
				assert.Equal(t, "key", key)
				return 1, true
			},
		})

		_, err := adapter.Get(context.Background(), "key")

		assert.NoError(t, err)
	})
}

func TestGetMulti(t *testing.T) {

	adapter := New(&stub{
		OnGet: func(key interface{}) (value interface{}, ok bool) {
			switch key {
			case "foo":
				return "test1", true
			case "bar":
				return "test2", true
			}
			return nil, false
		},
	})

	t.Run("return all requested values", func(t *testing.T) {
		actualItems, err := adapter.GetMulti(context.Background(), []interface{}{"foo", "bar"})
		assert.NoError(t, err)
		assert.EqualValues(t, map[interface{}]interface{}{
			"foo": "test1",
			"bar": "test2",
		}, actualItems)
	})

	t.Run("return part of requested values", func(t *testing.T) {
		actualItems, err := adapter.GetMulti(context.Background(), []interface{}{"foo", "missing"})
		assert.NoError(t, err)
		assert.EqualValues(t, map[interface{}]interface{}{
			"foo": "test1",
		}, actualItems)
	})

	t.Run("return nothing", func(t *testing.T) {
		actualItems, err := adapter.GetMulti(context.Background(), []interface{}{"missing1", "missing2"})
		assert.NoError(t, err)
		assert.Empty(t, actualItems)
	})
}

func TestSet(t *testing.T) {
	t.Run("WasCalled", func(t *testing.T) {
		wasCalled := false
		adapter := New(&stub{
			OnAdd: func(key, value interface{}) {
				wasCalled = true
			},
		})

		adapter.Set(context.Background(), "key", 1)

		assert.True(t, wasCalled)
	})

	t.Run("SameParams", func(t *testing.T) {
		adapter := New(&stub{
			OnAdd: func(key, value interface{}) {
				assert.Equal(t, "key", key)
				assert.Equal(t, 1, value)
			},
		})

		adapter.Set(context.Background(), "key", 1)
	})
}

func TestRemove(t *testing.T) {
	t.Run("WasCalled", func(t *testing.T) {
		wasCalled := false
		remove := func(key interface{}) (present bool) {
			wasCalled = true
			return
		}

		adapter := New(&stub{
			OnRemove: remove,
		})

		adapter.Remove(context.Background(), "key")

		assert.True(t, wasCalled)
	})

	t.Run("SameParams", func(t *testing.T) {
		remove := func(key interface{}) (present bool) {
			assert.Equal(t, "key", key)
			return
		}
		adapter := New(&stub{
			OnRemove: remove,
		})

		adapter.Remove(context.Background(), "key")
	})
}
