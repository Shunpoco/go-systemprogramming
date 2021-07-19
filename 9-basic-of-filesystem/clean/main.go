package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	// そのままクリーン
	fmt.Println(filepath.Clean("./path/filepath/../path.go"))

	// 絶対パスに整形
	abspath, _ := filepath.Abs("path/filepath/../path_unix.go")
	fmt.Println(abspath)

	// パスを相対パスに整形
	relpath, _ := filepath.Rel("/usr/local/go/src", "/usr/local/go/src/path/filepath/path.go")
	fmt.Println(relpath)

	// 環境変数の展開
	path := os.ExpandEnv("${PATH}")
	fmt.Println(path)

	// ~をホームディレクトリとして展開する
	my, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println(my.HomeDir)

	fmt.Println(cleanHomedir("~/hogehgoe/.././fuga/hoge.txt"))
}

func cleanHomedir(path string) string {
	if len(path) > 1 && path[0:2] == "~/" {
		my, err := user.Current()
		if err != nil {
			panic(err)
		}
		path = my.HomeDir + path[1:]
	}
	path = os.ExpandEnv(path)

	return filepath.Clean(path)
}
