package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		os.Exit(1)
	}

	filepath := os.Args[1]

	_, err := os.Stat(filepath)
	// if err == os.ErrNotExist { // これはmatchしない
	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("File %s doesn't exist", filepath)
		os.Exit(1)
	} else if err != nil {
		panic(err)
	}
}
