package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	info, err := os.Stat("hoge.txt")
	if os.IsNotExist(err) {
		fmt.Println("File does not exist")
		os.Exit(1)
	} else if err != nil {
		panic(err)
	}

	internalStat := info.Sys().(*syscall.Stat_t)

	fmt.Printf("%v\n", internalStat.Ino)
}
