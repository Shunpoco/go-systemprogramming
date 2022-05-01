package main

import (
	"fmt"
	"math"
)

// calc prime numbers like a generator function
func primeNumber() chan int {
	result := make(chan int)
	go func() {
		result <- 2
		for i := 3; i < 100; i += 1 {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j < l+1; j += 1 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i
			}
		}
		close(result)
	}()

	return result
}

func main() {
	pn := primeNumber()
	for n := range pn {
		fmt.Println(n)
	}
}
