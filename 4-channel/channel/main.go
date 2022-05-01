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
	task, ok := <-tasks
	fmt.Println(task)
	if ok {
		fmt.Println("Close!")
	}

	tasks <- "discord"
	// if you don't need to get data, we can discord it.
	<-tasks
}
