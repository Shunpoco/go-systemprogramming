package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path := os.Getenv("GOPATH")

	dir, name := filepath.Split(path)
	fmt.Printf("Dir: %s, Name: %s\n", dir, name)

	fragments := strings.Split(path, string(filepath.Separator))
	fmt.Printf("%v\n", fragments)

	fmt.Printf("%s\n", filepath.Base(path))
	fmt.Printf("%s\n", filepath.Dir(path))
	fmt.Printf("%s\n", filepath.Dir(filepath.Join(os.TempDir(), "temp.txt")))
}
