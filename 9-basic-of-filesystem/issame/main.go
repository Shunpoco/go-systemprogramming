package main

import (
	"fmt"
	"os"
)

func main() {
	info1, _ := os.Stat("hoge.txt")
	info2, _ := os.Stat("hoge_link.txt")

	if os.SameFile(info1, info2) {
		fmt.Println("Same!")
	}
}
