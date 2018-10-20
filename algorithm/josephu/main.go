package main

import "fmt"

//编写一个函数，构成单向的环形链表
type Boy struct {
	No   int  // 编号
	Next *Boy // 指向下一个小孩的指针
}

// 表示小孩的个数
// *Boy 返回该环形的链表的第一个小孩的指针
func AddBoy(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}
	// 判断
	if num < 1 {
		fmt.Println("num的值不对")
		return first
	}
	// 循环的构建这个环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first // 构造环形链表
		}
	}
	return first
}

func ShowBoy(first *Boy) {
	if first.Next == nil {
		fmt.Println("链表为空，没有小孩...")
		return
	}
	curBoy := first
	for {
		//
		fmt.Printf("小孩编号=%d ->", curBoy.No)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
}

func PlayGame(first *Boy, startNo int, countNum int) {
	if first.Next == nil {
		fmt.Println("空的链表，没有小孩")
		return
	}
	tail := first
	for {
		if tail.Next == first {
			// 说明tail到了最后的小孩
			break
		}
		tail = tail.Next
	}
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	for {
		for i := 1; i < countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号%d出圈 ->", first.No)
		// 删除first指向的小孩
		first = first.Next
		tail.Next = first
		if tail == first {
			break
		}
	}
}

func main() {
	first := AddBoy(5)
	ShowBoy(first)
}
