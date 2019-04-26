package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

/*
	如何从无序链表中移除重复项
	给定一个没有排序的链表，去掉其重复项，并保留原顺序
*/

//  顺序删除
func RemoveDup(head *gotype.LNode) {
	if head == nil || head.Next == nil {
		return
	}
	outerCur := head.Next
	var innerCur *gotype.LNode
	var innerPre *gotype.LNode
	for ; outerCur != nil; outerCur = outerCur.Next {
		for innerCur, innerPre = outerCur.Next, outerCur; innerCur != nil; {
			if outerCur.Data == innerCur.Data {
				innerPre.Next = innerCur.Next
				innerCur = innerCur.Next
			} else {
				innerPre = innerCur
				innerCur = innerCur.Next
			}
		}
	}
}

func main() {
	fmt.Println("删除重复结点")
}
