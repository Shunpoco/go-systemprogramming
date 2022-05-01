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
	done := make(chan struct{})
	go func() {
		fmt.Println("sub() is start!")
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
		done <- struct{}{}
	}()

	// Wait the finish
	<-done

	fmt.Println("all tasks are finished")
}
