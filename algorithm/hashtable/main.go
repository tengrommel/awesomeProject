package main

import "fmt"

// 定义emp
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

// 方法待定...
// 我们这里的EmpLink 不带表头，即第一个结点就存放雇员
type EmpLink struct {
	Head *Emp
}

// 添加员工的方法，保证添加时，编号是从小到大
func (e *EmpLink) Insert(emp *Emp) {
	cur := e.Head      // 这是辅助指针
	var pre *Emp = nil // 这是一个辅助指针 pre 在cur前面
	// 如果当前的EmpLink就是一个空链表
	if cur == nil {
		e.Head = emp
		return
	}
	// 如果不是一个空链表，给emp找到对应的位置并插入
	// 思路是让cur和emp比较，然后让pre 保持在cur前面
	for {
		if cur != nil {
			// 比较
			if cur.Id > emp.Id {
				// 找到位置
				break
			}
			pre = cur // 保证同步
			cur = cur.Next
		} else {
			break
		}
	}
	// 退出时，我们看下是否将emp添加到链表最后
	pre.Next = emp
	emp.Next = cur
}

// 定义hashtable，含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

func (e *EmpLink) ShowLink(no int) {
	if e.Head == nil {
		fmt.Printf("链表%d 为空\n", no)
		return
	}
	// 变量当前的链表，并显示数据
	cur := e.Head // 辅助的指针
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 名字=%s ->", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println() // 换行处理
}

// 给HashTable 编写Insert雇员的方法。
func (h *HashTable) Insert(emp *Emp) {
	// 使用散列函数确定将该雇员添加到哪个链表
	linkNo := h.HashFun(emp.Id)
	// 使用对应的链表添加
	h.LinkArr[linkNo].Insert(emp)
}

// 编写方法，显式hashtable所有的雇员
func (h *HashTable) ShowAll() {
	for i := 0; i < len(h.LinkArr); i++ {
		h.LinkArr[i].ShowLink(i)
	}
}

// 编写一个散列方法
func (h *HashTable) HashFun(id int) int {
	return id % 7 // 得到一个值，就是对于链表的下标
}

func main() {
	key := ""
	id := 0
	name := ""
	var hashtable HashTable
	for {
		fmt.Println("=============雇员系统菜单=============")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show 表示显示雇员")
		fmt.Println("find 表示查找雇员")
		fmt.Println("exit 表示退出系统")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员name")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashtable.Insert(emp)
		case "show":
			hashtable.ShowAll()
		case "exit":
		default:
			fmt.Println("输入错误")
		}
	}
}
