package main

import "os"

func main() {
	os.Link("hoge.txt", "hoge_hardln.txt")

	os.Symlink("hoge.txt", "hoge_symlink.txt")
}
