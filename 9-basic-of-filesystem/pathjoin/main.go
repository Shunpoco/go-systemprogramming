package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join(os.TempDir(), "temp.txt")
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Println("Does not exist!")
	}
}
