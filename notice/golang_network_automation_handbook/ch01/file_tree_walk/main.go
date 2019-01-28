package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
Recursively walking a file tree from a designated starting point can be used to find
and collect data or modify specify files in a directory structure.
The filepath.Walking() function, from the filepath package enables this functionality.
It accepts a root directory and a function to process the files and directories;
symbolic links are not followed.
*/

func main() {
	scan := func(path string, i os.FileInfo, _ error) error {
		fmt.Println(i.IsDir(), path)
		return nil
	}
	_ = filepath.Walk("/Users/tengzhou", scan)
}
