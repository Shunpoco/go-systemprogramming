package main

import (
	"fmt"
	"io"
	"os"
)

func open() {
	file, err := os.Create("textfile.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	io.WriteString(file, "New file content\n")
}

func read() {
	file, err := os.Open("textfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("Read file:")
	io.Copy(os.Stdout, file)
}

func main() {
	open()
	read()
}
