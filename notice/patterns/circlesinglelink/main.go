package main

import "fmt"

// 定义猫的结构体结点
type CatNode struct {
	no   int // 猫猫的编号
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	// 判断是不是添加第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head // 形成一个环状
		fmt.Println(newCatNode, "加入到环形的链表")
		return
	}
	// 定义一个临时变量，帮忙，找到环形的最后结点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	// 加入环形列表中
	temp.next = newCatNode
	newCatNode.next = head
	return
}

func ListCircleLink(head *CatNode) {
	fmt.Println("环形链表的具体情况如下")
	temp := head
	if temp.next == nil {
		fmt.Println("空空如也的环形链表...")
		return
	}
	for {
		fmt.Println("猫的信息为=", temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func main() {
	// 这里我们初始化一个环形链表的头结点
	head := &CatNode{}
	// 创建一只猫
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	InsertCatNode(head, cat1)
	ListCircleLink(head)
}
