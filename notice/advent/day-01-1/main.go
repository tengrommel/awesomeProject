package main

import (
	"fmt"
	"strings"
)

type Node struct {
	NodeMap map[string]Node
	Id      int64
}

type DatabaseItem struct {
	Name string
	Id   int64
}

var items = []DatabaseItem{
	{"a/b/c", 1},
	{"a/b/e", 2},
	{"a/d/f", 3},
	{"a", 4},
}

func main() {
	var node = Node{}
	for i := range items {
		currentNode := &node
		itemStringList := strings.Split(items[i].Name, "/")
		for index, item := range itemStringList {
			_, ok := currentNode.NodeMap[item]
			if !ok {
				currentNode.NodeMap = make(map[string]Node)
				currentNode.NodeMap[item] = Node{}
			}
			if index == len(itemStringList)-1 {
				currentNode.NodeMap[item] = Node{
					NodeMap: currentNode.NodeMap,
					Id:      items[i].Id,
				}
			} else {
				*currentNode = currentNode.NodeMap[item]
			}
		}
	}
	fmt.Println(node)
	fmt.Println("ok")
}
