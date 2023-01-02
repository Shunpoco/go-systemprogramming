package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:9999")

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())

			request, err := http.ReadRequest(
				bufio.NewReader(conn),
			)
			if err != nil {
				panic(err)
			}

			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(dump))

			response := http.Response{
				StatusCode: http.StatusAccepted,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body: io.NopCloser(
					strings.NewReader("Hello World\n"),
				),
			}
			response.Write(conn)
			conn.Close()
		}()
	}
}
