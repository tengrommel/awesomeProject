package main

import "github.com/isdamir/gotype"

// 快慢指针查询
func FindLastK(head *gotype.LNode, k int) *gotype.LNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head.Next
	fast := head.Next
	i := 0
	for i = 0; i < k && fast != nil; i++ {
		fast = fast.Next
	}
	if i < k {
		return nil
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func main() {

}
