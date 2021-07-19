package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]}\n", os.Args[0])
	}

	info, err := os.Stat(os.Args[1])
	if err == os.ErrNotExist {
		fmt.Printf("file not found: %s\n", os.Args[1])
	} else if err != nil {
		panic(err)
	}

	fmt.Println("FileInfo")
	fmt.Printf("	ファイル名: %v\n", info.Name())
	fmt.Printf("	サイズ: %v\n", info.Size())
	fmt.Printf("	変更日時: %v\n", info.ModTime())
	fmt.Println("Mode()")
	fmt.Printf("	ディレクトリ？: %v\n", info.Mode().IsDir())
	fmt.Printf("	読み書き可能な通常ファイル？: %v\n", info.Mode().IsRegular())
	fmt.Printf("	Unixのファイルアクセス権限ビット %o\n", info.Mode().Perm())
	fmt.Printf("	モードのテキスト表現: %v\n", info.Mode().String())

	internalStat := info.Sys().(*syscall.Stat_t)
	fmt.Println("OS固有のファイル属性")
	fmt.Printf("	デバイス番号: %v\n", internalStat.Dev)
	fmt.Printf("	inode番号: %v\n", internalStat.Ino)
	fmt.Printf("	ブロックサイズ: %v\n", internalStat.Blksize)
	fmt.Printf("	ブロック数: %v\n", internalStat.Blocks)
	fmt.Printf("	リンクされている数: %v\n", internalStat.Nlink)
	fmt.Printf("	最終アクセス日時: %v\n", internalStat.Atim)
	fmt.Printf("	属性変更日時: %v\n", internalStat.Ctim)
}
