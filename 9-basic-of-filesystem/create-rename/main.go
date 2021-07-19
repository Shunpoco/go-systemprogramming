package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// 新規作成
func open() {
	file, err := os.Create("textfile.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	io.WriteString(file, "New file content\n")
}

// 読み込み
func read() {
	file, err := os.Open("textfile.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	fmt.Println("Read file")
	io.Copy(os.Stdout, file)
}

// 追記モード
func append() {
	file, err := os.OpenFile("textfile.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	io.WriteString(file, "Append content\n")
}

// 削除
func remove() {
	err := os.Remove("textfile.txt")
	if err != nil {
		panic(err)
	}
}

// ディレクトリの作成
func mkdir(path string) {
	paths := strings.Split(path, "/")
	fmt.Println(paths)
	var err error
	if len(paths) > 1 {
		fmt.Println("greater than 2")
		err = os.MkdirAll(path, 0755)
	} else {
		err = os.Mkdir(path, 0755)
	}

	if err != nil {
		panic(err)
	}
}

// 指定したディレクトリを削除
func removeDir(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		panic(err)
	}
}

// fileの移動・リネーム
func rename() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	err = os.Mkdir("test", 0755)
	if err != nil {
		panic(err)
	}

	// rename
	err = os.Rename("test.txt", "new_test.txt")
	if err != nil {
		panic(err)
	}

	// 移動
	err = os.Rename("new_test.txt", "test/new_test.txt")
	if err != nil {
		panic(err)
	}

	// エラー
	err = os.Rename("test/new_test.txt", "test/")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	open()
	read()
	append()
	read()

	mkdir("test")
	removeDir("test")
	mkdir("test2/fuga")
	removeDir("test2")
	mkdir("test3/")
	removeDir("test3/")

	remove()

	rename()
}
