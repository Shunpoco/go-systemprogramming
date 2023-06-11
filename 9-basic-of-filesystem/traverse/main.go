package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var imageSuffix = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".webp": true,
	".gif":  true,
	".tiff": true,
	".eps":  true,
}

func main() {
	if len(os.Args) == 1 {
		fmt.Errorf(`Find images

Usage:
    %s [path to find]`,
			os.Args[0])
		os.Exit(1)
	}

	root := os.Args[1]
	root = filepath.Clean(root)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(info.Name()))
		if imageSuffix[ext] {
			rel, err := filepath.Rel(root, path)
			if err != nil {
				return nil
			}
			fmt.Printf("%s\n", rel)
		}
		return nil
	})
	if err != nil {
		fmt.Println(1, err)
	}
}
