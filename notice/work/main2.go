package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type node struct {
	Id        *int   `json:"Id"`
	Name      string `json:"name"`
	ChildNode []node `json:"child_node"`
}

type simple struct {
	id              int
	name            string
	contractpartyId int
}

func main() {
	f := []simple{
		{1, "a", 1},
		{2, "a/b/c/d", 1},
		{3, "a/c", 1},
	}
	reList := toTreeList(f)
	bt, _ := json.Marshal(reList)
	fmt.Println(string(bt))
}

func toList(simpleList []simple) []node {
	reList := make([]node, 0)
	for i := range simpleList {
		item := node{
			Id:   &simpleList[i].id,
			Name: simpleList[i].name,
		}
		reList = append(reList, item)
	}
	return reList
}

func toTreeList(simpleList []simple) []node {
	reList := make([]node, 0)
	for i := range simpleList {
		ll := strings.Split(simpleList[i].name, "/")
		nodeList := make([]node, 0)
		var re = node{Name: simpleList[i].name, ChildNode: nodeList}
		for j := 0; j < len(ll)-1; j++ {
			var b *int
			if j == 0 {
				b = &simpleList[i].id
			}
			re = nextNode(re, strings.Join(ll[:len(ll)-j-1], "/"), b)
		}
		reList = append(reList, re)
	}

	return reList
}

func nextNode(re node, name string, ii *int) node {
	if ii != nil {
		re.Id = ii
	}
	res := node{
		Name:      name,
		ChildNode: make([]node, 0),
	}
	res.ChildNode = append(res.ChildNode, re)
	return res
}
