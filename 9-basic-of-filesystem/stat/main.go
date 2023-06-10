package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]", os.Args[0])
		os.Exit(1)
	}

	info, err := os.Stat(os.Args[1])
	if err == os.ErrNotExist {
		fmt.Printf("file not found: %s\n", os.Args[1])
		os.Exit(1)
	} else if err != nil {
		panic(err)
	}

	fmt.Println("FileInfo")
	fmt.Printf("    Name: %v\n", info.Name())
	fmt.Printf("    Size: %v\n", info.Size())
	fmt.Printf("    Modified: %v\n", info.ModTime())
	fmt.Println("Mode()")
	fmt.Printf("    Is directory: %v\n", info.Mode().IsDir())
	fmt.Printf("    Is readable and writable: %v\n", info.Mode().IsRegular())
	fmt.Printf("    Permissions: %o\n", info.Mode().Perm())
	fmt.Printf("    Text: %v\n", info.Mode().String())
}
