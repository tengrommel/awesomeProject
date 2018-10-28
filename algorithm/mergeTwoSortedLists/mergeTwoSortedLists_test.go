package problem21

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	one []int
	two []int
}

type ans struct {
	one []int
}

func Test_Problem21(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		question{para{[]int{}, []int{1, 3, 5, 7}}, ans{[]int{1, 3, 5, 7}}},
	}
	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		ast.Equal(a.one, l2s(mergeTwoLists(s2l(p.one), s2l(p.two))))
	}
}

func l2s(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func s2l(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	res := &ListNode{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}
	return res
}
