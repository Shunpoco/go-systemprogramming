package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start sub()")
	// the channel in order to notice the end of goroutine.
	// we can use any types
	// Use struct, because it is 0byte.
	done := make(chan int)
	go func() {
		fmt.Println("sub() is start!")
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
		done <- 1
	}()

	// Wait the finish
	<-done

	fmt.Println("all tasks are finished")

	close(done)

	// After a channel is closed, we still can read data from the channel.
	// We get default value. In this case, int(0).
	fmt.Println(<-done)
}
