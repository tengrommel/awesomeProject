package main

import "fmt"

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	// 1、创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 // 黑子
	chessMap[2][3] = 2 // 蓝子
	// 2、输出看看原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	// 转成稀疏数组。想->算法
	// 思路
	/*
		1、遍历chessMap，如果我们发现有一个元素的值不为0，创建一个node结构体
		2、将其放入到对应的切片即可
	*/
	var sparseArr []ValNode
	// 标准的一个稀疏数组应该还含有一个 记录元素的二维数组的规模（行和列，默认值）

	valNode := ValNode{11, 11, 0}

	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				// 创建一个ValNode值节点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	// 输出稀疏数组
	fmt.Println("当前的稀疏数组")
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}

}
