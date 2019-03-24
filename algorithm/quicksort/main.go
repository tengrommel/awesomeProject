package main

import "fmt"

// 快速排序
func QuickSort(left int, right int, array *[6]int) {
	l := left
	r := right
	pivot := array[(right+left)/2]
	temp := 0
	// for循环的目的是将比pivot小的数放到左边
	// 比pivot大的数放在右边
	for l < r {
		// 先从pivot的左边找到大于等于pivot的值
		for array[l] < pivot {
			l++
		}
		// 从pivot的右边找到小于等于pivot的值
		for array[r] > pivot {
			r--
		}
		// 如果满足这个条件表明任务完成
		if l >= r {
			break
		}
		temp = array[l]
		array[l] = array[r]
		array[r] = temp
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		QuickSort(left, r, array)
	}
	if right > l {
		QuickSort(l, right, array)
	}
}

func main() {
	arr := [6]int{-9, 78, 0, 23, -567, 70}
	QuickSort(0, len(arr)-1, &arr)
	fmt.Println(arr)
}
