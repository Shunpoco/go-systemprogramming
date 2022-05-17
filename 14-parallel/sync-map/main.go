package main

import (
	"fmt"
	"sync"
)

func main() {
	smap := &sync.Map{}

	smap.Store("hello", "world")
	smap.Store(1, 2)

	smap.Delete("test")

	value, ok := smap.Load("hello")
	fmt.Printf("key=%v, value=%v, exists?=%v\n", "hello", value, ok)

	smap.Store("hoge", "fuga")

	smap.LoadOrStore(1, 10)
	smap.LoadOrStore("unko", "unko")

	smap.Range(func(key, value interface{}) bool {
		fmt.Printf("%v:%v\n", key, value)

		return true
	})
}
