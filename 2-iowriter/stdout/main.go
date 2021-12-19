package main

import "os"

func main() {
	// os.Stdout satisfies io.Writer.
	stdout := os.Stdout
	stdout.Write([]byte("os.Stdout example\n"))
}
