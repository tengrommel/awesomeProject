package main

import (
	"fmt"
	"strings"
)

type result struct {
	id        int
	name      string
	childNode []result
}

func main() {
	f := []struct {
		id       int
		name     string
		contract int
	}{
		{1, "a/b/c/d", 1},
	}
	ll := strings.Split(f[0].name, "/")
	nodeList := make([]result, 0)
	var re = result{name: f[0].name, childNode: nodeList}
	for i := 0; i < len(ll)-1; i++ {
		var b int
		if i == 0 {
			b = f[0].id
		}
		re = next(re, strings.Join(ll[:len(ll)-i-1], "/"), b)
	}
	fmt.Println(re)
}

func next(re result, name string, ii int) result {
	if ii != 0 {
		re.id = ii
	}
	res := result{
		name:      name,
		childNode: make([]result, 0),
	}
	res.childNode = append(res.childNode, re)
	return res
}
