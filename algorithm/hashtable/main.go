package main

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

// 给HashTable 编写Insert雇员的方法。
func (h *HashTable) Insert(emp *Emp) {
	// 使用散列函数确定将该雇员添加到哪个链表
	linkNo := h.HashFun(emp.Id)
	// 使用对应的链表添加
	h.LinkArr[linkNo].Insert(emp)
}

// 编写一个散列方法
func (h *HashTable) HashFun(id int) int {
	return id % 7 // 得到一个值，就是对于链表的下标
}

func main() {

}
