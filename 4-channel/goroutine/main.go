package main

import (
	"fmt"
	"time"
)

func sub() {
	fmt.Println("sub() is running")
	time.Sleep(time.Second)
	fmt.Println("Sub() is finished")
}

func main() {
	fmt.Println("Start sub()")
	go sub()
	time.Sleep(2 * time.Second)
}
