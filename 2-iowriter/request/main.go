package main

import (
	"net/http"
	"os"
)

func main() {
	// http.Request also satisfies io.Writer.
	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}

	request.Header.Set("X-TEST", "add header")
	request.Write(os.Stdout)
}
