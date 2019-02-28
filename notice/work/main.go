package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type result struct {
	Id        *int     `json:"Id"`
	Name      string   `json:"name"`
	ChildNode []result `json:"child_node"`
}

func main() {
	f := []struct {
		id       int
		name     string
		contract int
	}{
		{1, "a", 1},
		{2, "a/b/c/d", 1},
		{3, "a/c", 1},
	}
	reList := make([]result, 0)
	for i := range f {
		ll := strings.Split(f[i].name, "/")
		nodeList := make([]result, 0)
		var re = result{Name: f[i].name, ChildNode: nodeList}
		for j := 0; j < len(ll)-1; j++ {
			var b *int
			if j == 0 {
				b = &f[i].id
			}
			re = next(re, strings.Join(ll[:len(ll)-j-1], "/"), b)
		}
		reList = append(reList, re)
	}
	bt, _ := json.Marshal(reList)
	fmt.Println(string(bt))
	str1 := []string{"a", "b", "d"}
	str2 := []string{"a", "c", "d"}
	fmt.Println(maxNum(str1, str2))
}

func next(re result, name string, ii *int) result {
	if ii != nil {
		re.Id = ii
	}
	res := result{
		Name:      name,
		ChildNode: make([]result, 0),
	}
	res.ChildNode = append(res.ChildNode, re)
	return res
}

func maxNum(a []string, b []string) int {
	back := 0
	lessLen := len(a)
	if len(a) > len(b) {
		lessLen = len(b)
	}
	for i := 0; i < lessLen; i++ {
		if a[i] == b[i] {
			back = back + 1
		} else {
			break
		}
	}
	return back
}
