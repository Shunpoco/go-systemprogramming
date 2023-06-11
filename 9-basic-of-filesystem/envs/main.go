package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := os.ExpandEnv("${GOPATH}/hoge.go")
	fmt.Println(path)

	// Home dir
	homeDir, _ := os.UserHomeDir()
	fmt.Println(homeDir)

	fmt.Println(clean2("~/Document/hoge/fuga/../fuga/hoge.go"))
}

func clean2(path string) string {
	if len(path) > 1 && path[0:2] == "~/" {
		home, _ := os.UserHomeDir()
		path = home + path[1:]
	}
	path = os.ExpandEnv(path)

	return filepath.Clean(path)
}
