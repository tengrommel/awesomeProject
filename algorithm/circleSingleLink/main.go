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
		head.next = head // 形成一个环形
		fmt.Println(newCatNode, "加入到环形的链表")
		return
	}
	// 定义一个临时变量，帮忙，找到环形的最后这个结点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	// 加入到链表中
	temp.next = newCatNode
	newCatNode.next = head
}

// 删除一只猫
func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head
	if temp.next == nil {
		fmt.Println("这是一个空的环形链表，不能删除")
		return head
	}
	//如果只有一个结点
	if temp.next == head { // 只有一个结点
		temp.next = nil
		return head
	}
	flag := true
	for {
		if temp.next == head {
			break
		}

		if temp.no == id {
			if temp == head {
				head = head.next
			}
			// 找到了退出
			helper.next = temp.next
			fmt.Printf("猫猫=%d\n", id)
			flag = false
		}
		temp = temp.next
		helper = helper.next
	}
	if flag {
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("猫猫=%d\n", id)
		}
	}
	return head
}

//输出这个环形的链表
func ListCircleLink(head *CatNode) {
	fmt.Println("环形链表的具体情况如下")
	temp := head
	if temp.next == nil {
		fmt.Println("空空如也的环形链表...")
		return
	}
	for {
		fmt.Printf("猫的信息为=[id=%d name=%v] ->\n", temp.no, temp.name)
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
	cat2 := &CatNode{
		no:   2,
		name: "jim",
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	ListCircleLink(head)
}
