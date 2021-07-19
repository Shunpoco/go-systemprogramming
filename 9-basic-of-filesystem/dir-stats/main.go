package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open("/")
	if err != nil {
		panic(err)
	}

	fileInfos, err := dir.Readdir(-1) // 負の値を渡すとDir内全要素
	if err != nil {
		panic(err)
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			fmt.Printf("[Dir]	%s\n", fileInfo.Name())
		} else {
			fmt.Printf("[File]	%s\n", fileInfo.Name())
		}
	}
}
