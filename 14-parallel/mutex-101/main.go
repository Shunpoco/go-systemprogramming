package main

import (
	"fmt"
	"sync"
)

var id int

func generatedId(mutex *sync.Mutex) int {
	// Lock()/Unlock()をペアで呼び出してロックする
	mutex.Lock()
	defer mutex.Unlock()
	id++
	result := id
	return result
}

func main() {
	var mutex sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generatedId(&mutex))
		}()
	}
}
