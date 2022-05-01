package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func generator(ctx context.Context, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-ctx.Done():
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
	// A context created by context.WithDeadline is automatically closed if the time goes to the second value.
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	gen := generator(ctx, 1)

	wg.Add(1)

LOOP:
	for i := 0; i < 5; i++ {
		select {
		case result, ok := <-gen:
			if ok {
				fmt.Println(result)
			} else {
				fmt.Println("timeout")
				break LOOP
			}
		}
	}

	// We can also use cancel as an explicitly close
	cancel()

	wg.Wait()
}
