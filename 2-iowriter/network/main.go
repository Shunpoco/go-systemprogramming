package main

import (
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	// net.Dial returns net.Conn, which satisfies both io.Writer and io.Reader.
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}

	// net.Conn requests to a server.
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")

	// This is the method of io.Copy
	io.Copy(os.Stdout, conn)

	req, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	req.Write(conn)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponseWriter sample")
}
