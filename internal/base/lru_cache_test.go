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

		var size int
		fakeCache := base.LruCache{}

		cache := base.NewLruCache(size)
		result := cache.GetSize()
		expected := fakeCache.GetSize()

		assert.Equal(t, expected, result)
		assert.Nil(t, cache.Head)
		assert.Nil(t, cache.Tail)
	})

	t.Run(
		"Put first node at empty cache", func(t *testing.T) {
			t.Parallel()

			ptr := &base.Node{
				Key:   "1",
				Value: "1",
			}

			fake_cache := base.LruCache{
				Head: ptr,
				Tail: ptr,
			}

			size := 3
			cache := base.NewLruCache(size)
			cache.Put(ptr.Key, ptr.Value)

			assert.Equal(t, 1, cache.Len())

			assert.NotNil(t, cache.Head)
			assert.Equal(t, fake_cache.Head.Key, cache.Head.Key)
			assert.Equal(t, fake_cache.Head.Value, cache.Head.Value)

			assert.Equal(t, cache.Head, cache.Tail)

			assert.Nil(t, cache.Head.Next)
			assert.Nil(t, cache.Head.Prev)
			assert.Nil(t, cache.Tail.Next)
			assert.Nil(t, cache.Tail.Prev)
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

			expected := base.LruCache{
				Head: ptr1,
				Tail: ptr3,
			}

			size := 3
			cache := base.NewLruCache(size)
			cache.Put(ptr3.Key, ptr3.Value)
			cache.Put(ptr2.Key, ptr2.Value)
			cache.Put(ptr1.Key, ptr1.Value)

			assert.Equal(t, size, cache.Len())

			assert.Equal(t, expected.Head.Key, cache.Head.Key)
			assert.Equal(t, expected.Head.Value, cache.Head.Value)

			assert.Equal(t, expected.Tail.Key, cache.Tail.Key)
			assert.Equal(t, expected.Tail.Value, cache.Tail.Value)

			assert.Equal(t, "2", cache.Head.Next.Key)
			assert.Equal(t, "3", cache.Head.Next.Next.Key)

			assert.Equal(t, "2", cache.Tail.Prev.Key)
			assert.Equal(t, "1", cache.Tail.Prev.Prev.Key)
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

			expected := base.LruCache{
				Head: ptrD,
				Tail: ptrB,
			}
			expected.SetSize(size)

			assert.Equal(t, size, cache.Len())
			assert.Equal(t, ptrD.Key, cache.Head.Key)
			assert.Equal(t, ptrD.Value, cache.Head.Value)

			assert.Equal(t, ptrA.Key, cache.Head.Next.Key)
			assert.Equal(t, ptrB.Key, cache.Head.Next.Next.Key)
			assert.Nil(t, cache.Head.Next.Next.Next)

			deletedNode := cache.Get(ptrC.Key)
			assert.Nil(t, deletedNode)
		})

	/**
	befor: Head → ABC
	after delete C: Head → AB
	after insert C: Head → CAB
	*/
	t.Run(
		"Get node", func(t *testing.T) {
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

			fake_cache := base.LruCache{
				Head: ptrC,
				Tail: ptrB,
			}

			size := 3
			cache := base.NewLruCache(size)
			cache.Put(ptrC.Key, ptrC.Value)
			cache.Put(ptrB.Key, ptrB.Value)
			cache.Put(ptrA.Key, ptrA.Value)
			cache.Get(ptrC.Key)

			expected := fake_cache
			result := cache

			assert.Equal(t, expected.Head.Key, result.Head.Key)
			assert.Equal(t, expected.Tail.Key, result.Tail.Key)
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

			result := *cache.Get("3")

			assert.Equal(t, "new value", result)
			assert.Equal(t, 2, cache.Len())
		})

	t.Run(
		"If Invalid cache size - then default size 0", func(t *testing.T) {
			t.Parallel()

			invalidSize := -100
			cache := base.NewLruCache(invalidSize)

			assert.Equal(t, 0, cache.GetSize())

			cache.SetSize(invalidSize)
			assert.Equal(t, 0, cache.GetSize())
		})
}
