package main

import (
	"context"
	"errors"
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
				if err := ctx.Err(); errors.Is(err, context.Canceled) {
					fmt.Println("canceled")
				} else if errors.Is(err, context.DeadlineExceeded) {
					fmt.Println("deadline")
				}
				break LOOP
			case out <- num:
			}
			time.Sleep(2 * time.Second)
		}

		close(out)
		fmt.Println("generator closed")
	}()

	return out
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	gen := generator(ctx, 1)

	wg.Add(1)

	for i := 0; i < 5; i++ {
		result, ok := <-gen
		if ok {
			fmt.Println(result)
		} else {
			fmt.Println("timeout")
		}
	}

	cancel()

	wg.Wait()
}
