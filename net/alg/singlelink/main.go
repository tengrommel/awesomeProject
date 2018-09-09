package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode
}

// 给链表插入节点
// 编写第一种插入方式，在单链表最后加入。
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	// 思路
	// 1、先找到该链表的最后这个结点
	// 2、创建一个辅助结点[跑龙套，帮忙]
	temp := head
	for {
		if temp.next == nil { // 表示找到了最后
			break
		}
		temp = temp.next
	}
	//3、将newHeroNode加入到链表的最后
	temp.next = newHeroNode
}

func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no >= newHeroNode.no {
			break
		} else if temp.next.no == newHeroNode.no {
			flag = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Println("对不起，已经存在no=", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}
}

func ListHeroNode(head *HeroNode) {
	temp := head
	// 先判断该链表是不是空的链表
	if temp.next == nil {
		fmt.Println("空空如也...")
		return
	}
	for {
		fmt.Printf("[%d, %s, %s]", temp.next.no,
			temp.next.name, temp.next.nickname)
		// 判断是否链表后
		temp = temp.next
		if temp.next == nil {
			break
		}
		fmt.Print("-->")
	}
}

func main() {
	// 1、先创建一个头节点
	head := &HeroNode{}
	// 2、创建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "未知",
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}
	// 加入
	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)
	InsertHeroNode(head, hero3)
	ListHeroNode(head)
}
