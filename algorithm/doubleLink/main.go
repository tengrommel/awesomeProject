package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode
	next     *HeroNode
}

func insertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

// 显示链表的所有结点信息
// 这里仍然使用单向的链表显示方式
func ListHeroNode(head *HeroNode) {
	// 1、先找到该链表的最后这个结点
	temp := head
	// 先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空链表")
	}
	// 2、变量这个链表
	for {
		//if temp.next == nil {
		//}
		fmt.Printf("[%d , %s, %s]==>", temp.next.no,
			temp.next.name, temp.next.nickname)
		// 判断是否链表后
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func ListHeroNode2(head *HeroNode) {
	// 1、先找到该链表的最后这个结点
	temp := head
	// 先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空链表")
	}
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	// 2、变量这个链表
	for {
		//if temp.next == nil {
		//}
		fmt.Printf("[%d , %s, %s]==>", temp.no,
			temp.name, temp.nickname)
		// 判断是否链表头
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
}

func main() {
	head := &HeroNode{}
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}
	insertHeroNode(head, hero1)
	insertHeroNode(head, hero2)
	insertHeroNode(head, hero3)

	ListHeroNode(head)
	fmt.Println()
	ListHeroNode2(head)
}
