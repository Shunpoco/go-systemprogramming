package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	err := os.Mkdir("setting", 0755)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll("setting/myapp/networksettings", 0755)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("setting/myapp/networksettings/access.log")
	if err != nil {
		panic(err)
	}

	io.WriteString(file, "赤パジャマ青パジャマ黃パジャマ、隣の客はよく柿食う客だ、東京特許許可局")

	file.Close()

	// Truncate, UTF-8では全角は3バイトだよ
	err = os.Truncate("setting/myapp/networksettings/access.log", 9)

	// Open opens the file READ_ONLY mode. we have to use OpenFile to truncate the file.
	file, err = os.OpenFile("setting/myapp/networksettings/access.log", os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("A")
	io.Copy(os.Stdout, file)

	file.Seek(0, 0)
	err = file.Truncate(6)
	if err != nil {
		panic(err)
	}
	fmt.Println("B")
	io.Copy(os.Stdout, file)

	// Remove
	err = os.RemoveAll("setting/myapp")
	if err != nil {
		panic(err)
	}

	err = os.Remove("setting")
	if err != nil {
		panic(err)
	}
}
