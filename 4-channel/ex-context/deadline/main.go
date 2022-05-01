package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func generator(done chan struct{}, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-done:
				break LOOP
				// It pretends to take long time
				// case out <- num:
			}
		}
		close(out)
		fmt.Println("generator closed")
	}()

	return out
}

func main() {
	done := make(chan struct{})
	gen := generator(done, 1)
	deadlineChan := time.After(time.Second)

	wg.Add(1)

LOOP:
	for i := 0; i < 5; i++ {
		select {
		case result := <-gen:
			fmt.Println(result)
		case <-deadlineChan:
			fmt.Println("timeout")
			break LOOP
		}
	}
	close(done)

	wg.Wait()
}
