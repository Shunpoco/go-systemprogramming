package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("hoge")
		wg.Done()
	}()

	go func() {
		fmt.Println("fuga")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("終了")
}
