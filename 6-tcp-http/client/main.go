package main

import "net"

func main() {
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		panic(err)
	}
	println(conn)
}
