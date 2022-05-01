package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan int
	select {
	case m := <-c:
		func(int) {}(m)
	case <-time.After(2 * time.Second):
		fmt.Println("timed out")
	}
}
