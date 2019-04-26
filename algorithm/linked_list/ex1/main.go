package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

// 方法一：带头结点的逆序
func Reverse(node *gotype.LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var pre *gotype.LNode
	var cur *gotype.LNode
	next := node.Next
	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
	node.Next = pre
}

// 方法二：递归法
// 1、先逆序除第一个结点以外的子链表，接着把结点头添加到逆序的子链表的后面
// 2、递归

/*
	由于递归法也只需要对链表进行一次遍历，因此，算法的时间复杂度也为O(n)，其中，n为链表的长度。
	递归法的主要优点是：思路比较直观，容易理解，而且也不需要保存前驱结点的地址。
	缺点是：需要处理额外的压栈与弹栈操作。
*/

func RecursiveReverse(node *gotype.LNode) {
	firstNode := node.Next
	// 递归调用
	newHead := RecursiveReverseChild(firstNode)
	node.Next = newHead
}

func RecursiveReverseChild(node *gotype.LNode) *gotype.LNode {
	if node == nil || node.Next == nil {
		return node
	}
	newHead := RecursiveReverseChild(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}

// 方法三：插入法
// 插入法的思路：从链表的第二个结点开始，把遍历到的结点插入到头结点的后面，直到遍历结束。
func InsertReverse(node *gotype.LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var cur *gotype.LNode
	var next *gotype.LNode
	cur = node.Next.Next
	// 设置链表第一个结点为尾结点
	node.Next.Next = nil
	// 把遍历到的结点插入到头结点的后面
	for cur != nil {
		next = cur.Next
		cur.Next = node.Next
		node.Next = cur
		cur = next
	}
}

func main() {
	head := &gotype.LNode{}
	fmt.Println("就地逆序")
	gotype.CreateNode(head, 8)
	gotype.PrintNode("逆序前：", head)
	Reverse(head)
	gotype.PrintNode("逆序后：", head)
	fmt.Println("递归逆序")
	gotype.CreateNode(head, 8)
	gotype.PrintNode("逆序前：", head)
	RecursiveReverse(head)
	gotype.PrintNode("逆序后：", head)
}
