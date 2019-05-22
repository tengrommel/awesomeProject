package main

// KeyValue type
type KeyValue interface {
	LessThan(KeyValue) bool
	EqualTo(KeyValue) bool
}

// TreeNode class
type TreeNode struct {
	KeyValue KeyValue
	BalanceValue int
	LinkedNodes [2]*TreeNode
}

// opposite method
func opposite(nodeValue int) int {
	return 1 - nodeValue
}

// single rotation method
func singleRotation(rootNode *TreeNode, nodeValue int) *TreeNode {
	var saveNode *TreeNode
	saveNode = rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode.LinkedNodes[nodeValue]
	saveNode.LinkedNodes[nodeValue] = rootNode
	return saveNode
}

func main() {
	
}
