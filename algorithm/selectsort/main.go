package main

import "fmt"

// 编写函数selectSort完成排序
func SelectSort(arr *[5]int) {
	// 标准的访问方式
	for j := 0; j < len(arr)-1; j++ {
		// 先完成将第一个大值和arr[0] => 先易后难
		// 1、先假设 arr[0] 最大值
		max := arr[j]
		maxIndex := j
		// 2、遍历后面 1---[len(arr) - 1]比较
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] { // 找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}
		// 交换
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
		fmt.Printf("第%d次%v\n", j+1, *arr)
	}
}

func main() {
	// 定义一个数组
	arr := [5]int{10, 34, 19, 100, 80}
	fmt.Println(arr)
	SelectSort(&arr)
	fmt.Println(arr)
}
