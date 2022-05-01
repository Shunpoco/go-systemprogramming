package main

import (
	"fmt"
	"time"
)

func timer(d time.Duration) time.Time {
	var c chan int
	var r time.Time
	select {
	case m := <-c:
		func(int) {}(m)
	case d := <-time.After(2 * time.Second):
		r = d
	}
	return r
}

func main() {
	fmt.Println(timer(time.Second))
}
