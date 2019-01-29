package main

import (
	"fmt"
	"regexp"
)

/*
Using parentheses in the search pattern identifies individual match groups.
*/

func main() {
	data := `# host.domain.tld #`
	pattern := regexp.MustCompile(`(\w+).(\w+).(\w+)`)
	groups := pattern.FindAllStringSubmatch(data, -1)
	fmt.Printf("\n%q\n", groups)
	fmt.Printf("groups[0][0]: %s\n", groups[0][0])
	fmt.Printf("groups[0][1]: %s\n", groups[0][1])
	fmt.Printf("groups[0][2]: %s\n", groups[0][2])
	fmt.Printf("groups[0][3]: %s\n", groups[0][3])
}
