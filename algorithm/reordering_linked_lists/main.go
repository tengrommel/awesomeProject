package main

import "github.com/isdamir/gotype"

// 找出链表Head的中间结点，把链表从中间断成两个子链表
func findMiddleNode(head *gotype.LNode) *gotype.LNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast := head
	slow := head
	slowPre := head
	for fast != nil && fast.Next != nil {
		slowPre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	slowPre.Next = nil
	return slow
}

// 对不带头结点的单链表翻转
func reverse(head *gotype.LNode) *gotype.LNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *gotype.LNode
	var next *gotype.LNode
	for head != nil{
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func Reorder(head *gotype.LNode)  {
	if head==nil||head.Next==nil {
		return
	}
	cur1 := head.Next
	mid := findMiddleNode(head.Next)
	cur2 := reverse(mid)
	var tmp *gotype.LNode
	for cur1.Next != nil {
		tmp = cur1.Next
		cur1.Next = cur2
		cur1 = tmp
		tmp = cur2.Next
		cur2.Next = cur1
		cur2 = tmp
	}
	cur1.Next = cur2
}

func main() {
	
}
