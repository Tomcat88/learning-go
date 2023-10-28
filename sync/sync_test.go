package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing a counter three times leaves it a 3", func(t *testing.T) {
		wantedCount := 1000
		c := &Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < 1000; i++ {
			go func() {
				c.Increment()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, c, wantedCount)
	})
}

func assertCounter(t *testing.T, c *Counter, want int) {
	t.Helper()
	if c.Value() != want {
		t.Errorf("got %d but want %d", c.Value(), want)
	}
}
