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

func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	// 思路
	// 1、先找到适当的结点的位置
	// 2、创建一个辅助结点[跑龙套，帮忙]
	temp := head
	flag := true
	// 让插入结点的no，和temp的下一个结点的no比较
	for {
		if temp.next == nil { // 说明到链表的最后
			break
		} else if temp.next.no > newHeroNode.no {
			// 说明newHeroNode 就应该插入到temp后面
			break
		} else if temp.next.no == newHeroNode.no {
			// 说明我们的链表中已经有这个node，就不让插入
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("对不起，已经存在no= ", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}
}

// 双向链表删除一个结点
func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		// 找到了
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("要删除的id不存砸")
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
