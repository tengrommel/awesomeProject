package link_in_a_single_list

import "github.com/isdamir/gotype"

// 判断单链表是否有环
func IsLoop(head *gotype.LNode) *gotype.LNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head.Next
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}

func FindLoopNode(head *gotype.LNode, meetNode *gotype.LNode) *gotype.LNode {
	first := head.Next
	second := meetNode
	for first != second {
		first = first.Next
		second = second.Next
	}
	return first
}
