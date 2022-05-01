package main

import "fmt"

func main() {
	// without buffer
	// tasks := make(chan string)
	// with buffer
	tasks := make(chan string, 10)

	tasks <- "cmake .."
	tasks <- "cmake . --build Debug"

	task := <-tasks
	fmt.Println(task)
	// Get data and check close
	// the second variable shows that the channel is still open (true) or closed (false).
	task, ok := <-tasks
	fmt.Println(task)
	if ok {
		fmt.Println("Still open!")
	}

	tasks <- "discord"
	// if you don't need to get data, we can discord it.
	<-tasks
}
