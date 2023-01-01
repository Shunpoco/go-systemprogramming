package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			b, err := io.ReadAll(conn)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s\n", b)
			conn.Close()
		}()
	}
}
