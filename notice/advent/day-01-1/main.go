package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Node struct {
	NodeMap  map[string]*Node
	Id       int64
	FullName string
}

type DatabaseItem struct {
	Name string
	Id   int64
}

var items = []DatabaseItem{
	{"a/b/o", 0},
	{"a/b/c", 1},
	{"a/b/e", 2},
	{"a/d/f", 3},
	{"a", 4},
	{"f", 5},
}

func main() {
	var node = Node{}
	for i := range items {
		currentNode := &node
		itemStringList := strings.Split(items[i].Name, "/")
		for index, item := range itemStringList {
			_, ok := currentNode.NodeMap[item]
			if !ok {
				if currentNode.NodeMap == nil {
					currentNode.NodeMap = make(map[string]*Node)
				}
				currentNode.NodeMap[item] = &Node{}
			}
			if index == len(itemStringList)-1 {
				if currentNode.NodeMap == nil {
					currentNode.NodeMap[item] = &Node{
						Id:       items[i].Id,
						FullName: items[i].Name,
					}
				} else {
					currentNode.NodeMap[item].Id = items[i].Id
					currentNode.NodeMap[item].FullName = items[i].Name
				}
			} else {
				currentNode = currentNode.NodeMap[item]
			}
		}
	}
	//fmt.Println(node)
	bt, _ := json.Marshal(node)
	fmt.Println(string(bt))
	fmt.Println("ok")
}
