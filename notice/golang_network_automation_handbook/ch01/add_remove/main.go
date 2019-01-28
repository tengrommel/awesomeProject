package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var file *os.File
	dpath := "./dir"
	fname := "file.txt"
	_ = os.MkdirAll(dpath, 0777)
	fpath := filepath.Join(dpath, fname)
	fmt.Println(fpath)
	defer file.Close()
	file, _ = os.Create(fpath)
	_ = os.Remove(fpath)
	_ = os.RemoveAll(dpath)
}
