package sync

import (
	"sync"
	"testing"
)

func TestCounterMutex(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounterMutex()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 10000
		counter := NewCounterMutex()

		var wg sync.WaitGroup
		wg.Add(wantedCount)
		for range wantedCount {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func BenchmarkCounterMutex(b *testing.B) {
	counter := NewCounterMutex()
	for b.Loop() {
		counter.Inc()
	}
}

func TestCounterAtomic(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounterAtomic()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 10000
		counter := NewCounterMutex()

		var wg sync.WaitGroup
		wg.Add(wantedCount)
		for range wantedCount {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func BenchmarkCounterAtomic(b *testing.B) {
	counter := NewCounterMutex()
	for b.Loop() {
		counter.Inc()
	}
}

func assertCounter(t testing.TB, got Counter, want int) {
	t.Helper()

	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
