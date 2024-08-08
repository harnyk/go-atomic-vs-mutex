package avm

import (
	"sync"
	"sync/atomic"
)

const numGoroutines = 4

func SynchronizeWithMutexes(times uint) int64 {
	var value int64
	mu := sync.Mutex{}

	wg := sync.WaitGroup{}

	change := func(delta int64) {
		mu.Lock()
		value += delta
		mu.Unlock()
	}

	grind := func(times uint, delta int64) {
		for i := uint(0); i < times; i++ {
			change(delta)
		}
		wg.Done()
	}

	for i := 0; i < numGoroutines/2; i++ {
		wg.Add(1)
		go grind(times, 1)
		wg.Add(1)
		go grind(times, -1)
	}
	wg.Wait()

	return value
}

func SynchronizeWithAtomic(times uint) int64 {
	var value int64

	wg := sync.WaitGroup{}

	change := func(delta int64) {
		atomic.AddInt64(&value, delta)
	}

	grind := func(times uint, delta int64) {
		for i := uint(0); i < times; i++ {
			change(delta)
		}
		wg.Done()
	}

	for i := 0; i < numGoroutines/2; i++ {
		wg.Add(1)
		go grind(times, 1)
		wg.Add(1)
		go grind(times, -1)
	}
	wg.Wait()

	return value
}
