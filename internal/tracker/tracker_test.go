package tracker_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"job4j.ru/go-lang-base/internal/tracker"
)

func Test_tracker(t *testing.T) {
	t.Parallel()

	t.Run("check link leak", func(t *testing.T) {
		t.Parallel()

		instance := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}
		_, err := instance.AddItem(item)
		require.NoError(t, err, "addedItem should not return error")

		res := instance.GetItems()
		res[0].Name = "Second Item"

		assert.Equal(t,
			[]tracker.Item{item},
			instance.GetItems(),
		)
	})

	t.Run("error update - not found", func(t *testing.T) {
		t.Parallel()

		instance := tracker.NewTracker()
		item := tracker.Item{
			ID: "1",
		}

		err := instance.UpdateItem(item.ID, "new")
		assert.ErrorIs(t, err, tracker.ErrNotFound)
	})

	t.Run("error add - item has already exist", func(t *testing.T) {
		t.Parallel()

		instance := tracker.NewTracker()
		item := tracker.Item{
			ID: "1",
		}

		_, err1 := instance.AddItem(item)

		require.NoError(t, err1, "addedItem should not return error")

		_, err := instance.AddItem(item)

		assert.ErrorIs(t, err, tracker.ErrHasAlreadyExist)
	})
}
