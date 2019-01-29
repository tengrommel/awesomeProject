package main

import (
	"fmt"
	"regexp"
)

func main() {
	data := `# host.domain.tld #`
	pattern := regexp.MustCompile(`\w+.\w+.\w+`)
	result := pattern.FindString(data)
	fmt.Println(result)
}
