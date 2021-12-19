package main

import "os"

func main() {
	// os.Create returns a *os.File and an error, if any.
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// *os.File has Write method, so this satisfies io.Writer.
	file.Write([]byte("os.FIle example\n"))
}
