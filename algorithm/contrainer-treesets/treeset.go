package main

import "sync"

// TreeSet class
type TreeSet struct {
	bst *BinarySearchTree
}

// TreeNode class
type TreeNode struct {
	key       int
	value     int
	leftNode  *TreeNode //left
	rightNode *TreeNode //right
}

// BinarySearchTree class
type BinarySearchTree struct {
	rootNode *TreeNode
	lock     sync.RWMutex
}


func main() {
	
}
