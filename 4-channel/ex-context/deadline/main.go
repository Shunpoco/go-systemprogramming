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
	// We can check whether a context has a timeout or not.
	ctx := context.Background()
	// The second output of context.Deadline() shows the context has timeout(true) or not(false)
	fmt.Println(ctx.Deadline())

	fmt.Println(time.Now())
	// // A context created by context.WithDeadline is automatically closed if the time goes to the second value.
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	// context.WithTimeout treats time.Duration as its second variable.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	fmt.Println(ctx.Deadline())
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
	// In this case, cancel() does nothing.
	// The official document says that we should call cancel() explicitlily even if we set timeout.
	cancel()

	wg.Wait()
}
