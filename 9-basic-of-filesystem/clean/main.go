package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	messyPath := "./path/filepath/../path.go"

	fmt.Println(filepath.Clean(messyPath))

	absPath, _ := filepath.Abs("./main.go")
	fmt.Println(absPath)

	relpath, _ := filepath.Rel("/usr/local/go/src", "/usr/local/go/src/path/filepath/path.go")
	fmt.Println(relpath)

	// 特に探索をするわけでもないので、 ./main.goと同様の結果になる
	absPath, err := filepath.Abs("hoge.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(absPath)
}
