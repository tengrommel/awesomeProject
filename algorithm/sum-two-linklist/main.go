package main

import "github.com/isdamir/gotype"

func Add(h1 *gotype.LNode, h2 *gotype.LNode) *gotype.LNode {
	if h1 == nil || h1.Next == nil {
		return h2
	}
	if h2 == nil || h2.Next == nil {
		return h1
	}
	c := 0
	sum := 0
	p1 := h1.Next
	p2 := h2.Next
	resultHead := &gotype.LNode{}
	p := resultHead
	for p1 != nil && p2 != nil {
		p.Next = &gotype.LNode{}
		sum = p1.Data.(int) + p2.Data.(int) + c
		p.Next.Data = sum % 10
		c = sum / 10
		p = p.Next
		p1 = p1.Next
		p2 = p2.Next
	}
	// 链表h2比h1长，接下来只需要考虑h2剩余的值
	if p1 != nil {
		for p2 != nil {
			p.Next = &gotype.LNode{}
			sum = p2.Data.(int) + c
			p.Next.Data = sum % 10
			c = sum / 10
			p = p.Next
			p2 = p2.Next
		}
	}
	// 相反
	if p2 != nil {
		for p1 != nil {
			p.Next = &gotype.LNode{}
			sum = p1.Data.(int) + c
			p.Next.Data = sum % 10
			c = sum / 10
			p = p.Next
			p1 = p1.Next
		}
	}
	if c == 1 {
		p.Next = &gotype.LNode{}
		p.Next.Data = 1
	}
	return resultHead
}

func main() {

}
