package main

import "github.com/isdamir/gotype"

func RemoveDup(head *gotype.LNode) {
	if head == nil || head.Next == nil {
		return
	}
	outerCur := head.Next
	var innerCur *gotype.LNode // 用于外层循环，指向链表第一个结点
	var innerPre *gotype.LNode // 用于内层循环遍历outerCur后面的结点
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

func removeDupRecursionChild(head *gotype.LNode) *gotype.LNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pointer *gotype.LNode
	cur := head
	// 对以head.Next为首的子链表删除重复的结点
	head.Next = removeDupRecursionChild(head.Next)
	pointer = head.Next
	// 找出以以head.Next为首的子链表中与head结点相同的结点并删除
	for pointer != nil {
		if head.Data == pointer.Data {
			cur.Next = pointer.Next
			pointer = cur.Next
		} else {
			pointer = pointer.Next
			cur = cur.Next
		}
	}
	return head
}

func RemoveDupRecursion(head *gotype.LNode) {
	if head == nil {
		return
	}
	head.Next = removeDupRecursionChild(head.Next)
}

func main() {

}
