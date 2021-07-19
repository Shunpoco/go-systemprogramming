package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	file1, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	file2, err := os.Create("test2.txt")
	if err != nil {
		panic(err)
	}

	info1, err := file1.Stat()
	if err != nil {
		panic(err)
	}

	info2, err := file2.Stat()
	if err != nil {
		panic(err)
	}

	if os.SameFile(info1, info1) {
		fmt.Println("おなじ〜")
	}
	if !os.SameFile(info1, info2) {
		fmt.Println("ちがう〜")
	}

	// ファイルモードの変更
	file1.Chmod(0644)
	// ファイルオーナーの変更
	file2.Chown(os.Getuid(), os.Getgid())

	// ファイルの最終アクセス日時、変更日時を変更
	os.Chtimes("test.txt", time.Now(), time.Now())

	// ハードリンク
	os.Link("test.txt", "test-link.txt")

	// シンボリックリンク
	os.Symlink("test2.txt", "test2-sym.txt")

	// シンボリックリンクのリンク先を取得
	link, err := os.Readlink("test2-sym.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(link)

	file1.Close()
	file2.Close()

	// リンクの削除
	err = os.Remove("test-link.txt")
	if err != nil {
		panic(err)
	}

	err = os.Remove("test2-sym.txt")
	if err != nil {
		panic(err)
	}

	err = os.Remove("test.txt")
	if err != nil {
		panic(err)
	}
	err = os.Remove("test2.txt")
	if err != nil {
		panic(err)
	}
}
