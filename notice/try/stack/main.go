package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type node struct {
	Name  string
	Child []node
}

func main() {
	nodeList := make([]node, 0)
	str := []string{"a", "a/b", "a/b/c"}
	child := make([]node, 0)
	nodeList = append(nodeList, dd(str, 0, child))
	bt, _ := json.Marshal(nodeList)
	fmt.Println(string(bt))
}

func dd(ss []string, i int, child []node) (n node) {
	if i > len(ss)-1 {
		return
	}
	if len(strings.Split(ss[i], "/")) != 0 {
		n = node{
			Name:  ss[i],
			Child: child,
		}
		child = append(child, dd(ss, i+1, child))
	}

	return
}
