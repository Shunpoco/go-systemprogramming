package main

import "net"

func main() {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}
	println(conn)
}
