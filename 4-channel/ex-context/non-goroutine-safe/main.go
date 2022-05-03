package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	slice := make([]int, 0)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			slice = append(slice, i)
		}(i)
	}

	wg.Wait()

	// This code does not guarantee the output is 10.
	// Because slice is not `goroutine-safe`!
	fmt.Println(len(slice))
}
