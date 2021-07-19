package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Printf("Temp File Path: %s\n", filepath.Join(os.TempDir(), "temp.txt"))

	dir, name := filepath.Split(os.Getenv("PATH"))
	fmt.Printf("Dir: %s, Name: %s\n", dir, name)

	fragments := strings.Split("/a/b/c.txt", string(filepath.Separator))

	fmt.Println(fragments)

	base := filepath.Base("/a/b/c.txt")
	dir = filepath.Dir("/a/b/c.txt")
	ext := filepath.Ext("/a/b/c.txt")

	fmt.Println(base)
	fmt.Println(dir)
	fmt.Println(ext)
}
