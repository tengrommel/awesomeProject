package main

import "fmt"

func main() {
	fileExt := map[string]string{
		"Golang": ".go",
		"C++":    ".cpp",
		"Java":   ".java",
		"Python": ".py",
	}
	fmt.Println(fileExt)
	fmt.Println(len(fileExt))
	if ext, ok := fileExt["Java"]; ok {
		fmt.Println(ext, ok)
	}
	if ext, ok := fileExt["Java"]; ok {
		fmt.Println(ext, ok)
	} else {
		fmt.Println("Not found!")
	}
}
