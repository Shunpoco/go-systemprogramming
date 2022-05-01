package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("start sub()")

	// A context in order to receive a finish
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("sub() is finished")
		time.Sleep(time.Second)
		// Send a finish
		cancel()
	}()

	// Wait a finish
	<-ctx.Done()

	fmt.Println("All tasks are finished!")
}
