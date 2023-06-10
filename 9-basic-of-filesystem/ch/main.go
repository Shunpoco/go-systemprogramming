package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	os.Chmod("hoge.txt", 0655)
	os.Chown("hoge.txt", os.Getuid(), os.Getgid())
	fmt.Println(os.Getuid())
	fmt.Println(os.Getgid())
	os.Chtimes("hoge.txt", time.Now(), time.Now())
}
