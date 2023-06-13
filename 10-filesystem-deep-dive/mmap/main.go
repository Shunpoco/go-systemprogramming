package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/edsrzf/mmap-go"
)

func main() {
	var testData = []byte("0123456789ABCDEF")
	var testPath = filepath.Join(os.TempDir(), "testdata")
	err := ioutil.WriteFile(testPath, testData, 0644)
	if err != nil {
		panic(err)
	}

	// mapping
	f, err := os.OpenFile(testPath, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m, err := mmap.Map(f, mmap.COPY, 0)
	if err != nil {
		panic(err)
	}
	defer m.Unmap()

	m[9] = 'X'
	fmt.Printf("mmap:     %s\n", m)
	m.Flush()

	fileData, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Original: %s\n", testData)
	fmt.Printf("mmap:     %s\n", m)
	fmt.Printf("file:     %s\n", fileData)
}
