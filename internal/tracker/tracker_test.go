package tracker_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		instance.AddItem(item)

		res := instance.GetItems()
		res[0].Name = "Second Item"

		assert.Equal(t,
			[]tracker.Item{item},
			instance.GetItems(),
		)
	})
}
