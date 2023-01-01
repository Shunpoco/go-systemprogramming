package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(conn, "hoge")
	conn.Close()

	conn, err = net.Dial("tcp", "localhost:9999")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(conn, "fuga")
	conn.Close()
}
