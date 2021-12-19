package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	// bytes.Buffer satisfies io.Writer, it can keep bytes.
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())

	// buffer also has WriteString() method. But we can't use the method other structs which satisfy io.Writer.
	buffer.WriteString("hoge hoge\n")
	fmt.Println(buffer.String())

	// instead of using bytes.Buffer.WriteString, we can use io.WriteString.
	io.WriteString(&buffer, "fuga fuga\n")
	fmt.Println(buffer.String())
}
