package main

import (
	"fmt"
	"io/ioutil"
)

/*
Listing directory contents within a specific directory can be implemented using one of the I/O utility functions from the ioutil package.
By specifying the directory name in the ioutil.ReadDir() function,
it converniently returns a list of os.FileInfo objects, as seen in the previous section.
*/

func main() {
	files, _ := ioutil.ReadDir(".")
	for _, file := range files {
		fmt.Println(file.Name(), file.ModTime())
	}
}
