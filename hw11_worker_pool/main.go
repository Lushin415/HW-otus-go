package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("Goroutine", i, "start")
			defer wg.Done()
			mutex.Lock()
			count++
			fmt.Println(count)
			mutex.Unlock()
			fmt.Println("Goroutine", i, "end")
		}(i)
	}

	wg.Wait()
	fmt.Println("Done")
}
