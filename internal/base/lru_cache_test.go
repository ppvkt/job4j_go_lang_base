package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_lru_cache(t *testing.T) {
	t.Parallel()

	t.Run("Check if empty cache was created", func(t *testing.T) {
		t.Parallel()

		size := 3
		cache := base.NewLruCache(size)

		assert.Equal(t, 0, cache.Len())
		assert.Equal(t, 3, cache.GetSize())
	})

	t.Run(
		"Put first node at empty cache", func(t *testing.T) {
			t.Parallel()

			size := 3
			cache := base.NewLruCache(size)

			ok := cache.Put("1", "1")
			assert.True(t, ok)
			assert.Equal(t, 1, cache.Len())

			node := cache.Get("1")
			assert.NotNil(t, node)
			assert.Equal(t, "1", *node)
		})

	t.Run(
		"Put nodes until cache has place", func(t *testing.T) {
			t.Parallel()

			ptr1 := &base.Node{
				Key:   "1",
				Value: "1",
			}

			ptr2 := &base.Node{
				Key:   "2",
				Value: "2",
			}

			ptr3 := &base.Node{
				Key:   "3",
				Value: "3",
			}

			size := 3
			cache := base.NewLruCache(size)
			cache.Put(ptr3.Key, ptr3.Value)
			cache.Put(ptr2.Key, ptr2.Value)
			cache.Put(ptr1.Key, ptr1.Value)

			assert.Equal(t, size, cache.Len())

			for _, key := range []string{ptr1.Key, ptr2.Key, ptr3.Key} {
				val := cache.Get(key)
				assert.NotNil(t, val)
				assert.Equal(t, key, *val)
			}
		})

	/**
	befor: Head → ABC
	after put D: Head → DAB
	*/
	t.Run(
		"Put nodes more that cache has place",
		func(t *testing.T) {
			t.Parallel()

			ptrA := &base.Node{
				Key:   "1",
				Value: "1",
			}
			ptrB := &base.Node{
				Key:   "2",
				Value: "2",
			}
			ptrC := &base.Node{
				Key:   "3",
				Value: "3",
			}
			ptrD := &base.Node{
				Key:   "4",
				Value: "4",
			}

			size := 3
			cache := base.NewLruCache(size)
			cache.Put(ptrC.Key, ptrC.Value)
			cache.Put(ptrB.Key, ptrB.Value)
			cache.Put(ptrA.Key, ptrA.Value)
			cache.Put(ptrD.Key, ptrD.Value)

			assert.Equal(t, size, cache.Len())

			deletedNode := cache.Get(ptrC.Key)
			assert.Nil(t, deletedNode)

			value := cache.Get(ptrD.Key)
			assert.NotNil(t, value)
			assert.Equal(t, ptrD.Value, *value)
		})

	/**
	  befor: Head → 123
	  after get 3: Head → 312
	  after put 4: Head → 431
	*/
	t.Run("Get node moves it to the front", func(t *testing.T) {
		t.Parallel()

		size := 3
		cache := base.NewLruCache(size)

		cache.Put("3", "3")
		cache.Put("2", "2")
		cache.Put("1", "1")

		cache.Get("3")

		cache.Put("4", "4")

		assert.NotNil(t, cache.Get("3"), "3 should still be in cache after Get")
		assert.Nil(t, cache.Get("2"), "2 shouldn't still be in cache")
		assert.Equal(t, size, cache.Len())
	})

	t.Run(
		"Get invalid node - nil", func(t *testing.T) {
			t.Parallel()

			size := 3
			cache := base.NewLruCache(size)
			cache.Put("3", "3")
			cache.Put("2", "2")
			cache.Put("1", "1")

			result := cache.Get("4")

			assert.Nil(t, result)
		})

	t.Run(
		"Put node with the same key", func(t *testing.T) {
			t.Parallel()

			size := 3
			cache := base.NewLruCache(size)
			cache.Put("3", "3")
			cache.Put("2", "2")
			cache.Put("3", "new value")

			result := cache.Get("3")
			assert.NotNil(t, result)
			assert.Equal(t, "new value", *result)
			assert.Equal(t, 2, cache.Len())
		})

	t.Run(
		"If Invalid cache size - then panic", func(t *testing.T) {
			t.Parallel()

			invalidSize := -100

			assert.Panics(t, func() {
				base.NewLruCache(invalidSize)
			})
		})

	t.Run("Cache with size 0 should not store any items", func(t *testing.T) {
		t.Parallel()

		cache := base.NewLruCache(0)

		assert.Equal(t, 0, cache.GetSize())

		ok := cache.Put("1", "1")
		assert.False(t, ok)
		assert.Equal(t, 0, cache.Len())

		assert.Nil(t, cache.Get("1"))
	})
}
