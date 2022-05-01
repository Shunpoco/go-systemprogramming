package main

import (
	"fmt"
	"time"
)

func evenNumber() chan int {
	result := make(chan int)
	go func() {
		time.Sleep(time.Second)
		result <- 2
		for i := 3; i < 15; i++ {
			if i%2 == 0 {
				result <- i
			}
		}
		close(result)
	}()
	return result
}

func oddNumber() chan int {
	result := make(chan int)
	go func() {
		time.Sleep(time.Second)
		result <- 1
		for i := 2; i < 15; i++ {
			if i%2 == 1 {
				result <- i
			}
		}
		close(result)
	}()
	return result
}

func main() {
	e := evenNumber()
	o := oddNumber()

	// does not finish this loop
	for {
		select {
		case even := <-e:
			fmt.Println("Even: ", even)
		case odd := <-o:
			fmt.Println("Odd: ", odd)
		default:
			fmt.Println("Data is not coming.")
			time.Sleep(time.Second)
			break
		}
	}
}
