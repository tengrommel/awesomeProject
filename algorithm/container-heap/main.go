package main

import (
	"container/heap"
	"fmt"
)

type IntegerHeap []int

func (iHeap *IntegerHeap)Len() int {
	return len(*iHeap)
}

func (iHeap *IntegerHeap)Less(i, j int) bool {
	return (*iHeap)[i] < (*iHeap)[j]
}

func (iHeap *IntegerHeap)Swap(i, j int)  {
	(*iHeap)[i], (*iHeap)[j] = (*iHeap)[j], (*iHeap)[i]
}

// IntegerHeap method -pushes the item
func (iHeap *IntegerHeap)Push(heapIntElement interface{})  {
	*iHeap = append(*iHeap, heapIntElement.(int))
}

// IntegerHeap method -pops the item from the heap
func (iHeap *IntegerHeap)Pop() interface{} {
	var n int
	var x1 int
	var previous IntegerHeap = *iHeap
	n = len(previous)
	x1 = previous[n-1]
	*iHeap = previous[0: n-1]
	return x1
}

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("minimum: %d\n", (*intHeap)[3])
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}
