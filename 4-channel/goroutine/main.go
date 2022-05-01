package main

import (
	"fmt"
	"time"
)

func a() {
	fmt.Println("A!")
}

func main() {
	// Standard function call
	a()

	// Define an anonymous function and call it
	func() {
		fmt.Println("This is an anonymous function!")
	}()

	// Create an another goroutine and call a
	// probably it does not print anything
	// because main() is finished before this goroutine is finished.
	go a()

	// Define an anonymous function, then call it on an another goroutine
	go func() {
		// if does not print anything
		fmt.Println("This is an anonymous function 2!")
	}()

	// It's a very naive solution: just wait a second.
	// Goroutines may be finished.
	time.Sleep(100000)
}
