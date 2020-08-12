package syncv1

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Increase 1 twice", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 2)
	})

	t.Run("Concurrent 1000 Inc", func(t *testing.T) {
		toInc := 10000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(toInc)

		for i := 0; i < toInc; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, counter, toInc)
	})
}

func assertCounter(t *testing.T, counter *Counter, want int) {
	if counter.Value() != want {
		t.Errorf("got %d, want %d\n", counter.Value(), want)
	}
}
