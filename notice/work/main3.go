package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Mod struct {
	Id   int64
	Name string
}

type Node struct {
	Id       int64   `json:"id"`
	NodeName string  `json:"nodeName"`
	FullName string  `json:"fullName"`
	Children []*Node `json:"children"`
}

func nodeFindByName(nodeList []*Node, name string) *Node {
	for i := range nodeList {
		if nodeList[i].NodeName == name {
			return nodeList[i]
		}
	}
	return nil
}

func main() {
	var node = Node{}
	var exampleList = []Mod{{1, "a"}, {2, "a/b"}, {3, "a/b/c"}, {4, "b"}, {5, "b/a"}}
	for i := range exampleList {
		currentNode := &node
		stringList := strings.Split(exampleList[i].Name, "/")
		for index := range stringList {
			if currentNode.Children == nil {
				currentNode.Children = make([]*Node, 0)
			}
			n := nodeFindByName(currentNode.Children, stringList[index])
			if n == nil {
				currentNode.Children = append(currentNode.Children, &Node{
					Id:       exampleList[i].Id,
					NodeName: stringList[index],
					FullName: exampleList[i].Name,
				})
				break
			} else {
				currentNode = n
			}
		}
	}
	bt, _ := json.Marshal(node)
	fmt.Println(string(bt))
}
