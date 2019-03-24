package main

import "fmt"

func InsertSort(arr *[5]int) {
	for j := 1; j < len(arr); j++ {
		// 完成第一次，给第二个元素找到合适的位置并插入
		insertVal := arr[j]
		insertIndex := j - 1 // 下标
		// 从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] // 数据后移
			insertIndex--
		}
		// 插入
		if insertIndex+1 != j {
			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("第%d次插入后的结果 %v\n", j, *arr)
	}
}

func main() {
	arr := [5]int{23, 0, 12, 56, 34}
	InsertSort(&arr)
	fmt.Println("main 函数")
	fmt.Println(arr)
}
