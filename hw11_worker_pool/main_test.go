package main_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcurrentIncrement(t *testing.T) {
	var count int
	var mutex sync.Mutex
	var wg sync.WaitGroup

	numGoroutines := 10
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			mutex.Lock()
			count++
			mutex.Unlock()
		}()
	}

	wg.Wait()

	assert.Equal(t, numGoroutines, count, "Счетчик должен быть равен количеству горутин")
}
