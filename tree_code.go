package main

import (
	"fmt"
)

type BTreeNode struct {
	data  int
	left  *BTreeNode
	right *BTreeNode
}

func isEqualTo(root1 *BTreeNode, root2 *BTreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	} else if root1 == nil && root2 != nil || root1 != nil && root2 == nil {
		return false
	} else if root1.data != root2.data {
		return false
	} else {
		return isEqualTo(root1.left, root2.left) && isEqualTo(root1.right, root2.right)
	}
}

func main() {
	root1 := &BTreeNode{1, &BTreeNode{2, nil, nil}, &BTreeNode{3, nil, nil}}
	root2 := &BTreeNode{1, &BTreeNode{2, nil, nil}, &BTreeNode{3, nil, &BTreeNode{1, nil, nil}}}
	fmt.Println(isEqualTo(root1, root2))
}
