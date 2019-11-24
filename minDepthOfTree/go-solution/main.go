package main

import (
	"fmt"
)

func minDepth(root *TreeNode) int {
	var q, newQ [](TreeNode) // keep track of the nodes
	// put first node in
	if root == nil {
		return 0
	}
	q = append(q, *root)
	height := 0
	for {
		// dequeue all
		for _, node := range q {
			if node.Left == nil && node.Right == nil {
				// this node doesn't have anything
				return height + 1
			}
			if node.Left != nil {
				// has a left, put to queue
				newQ = append(newQ, *node.Left)
			}
			if node.Right != nil {
				// has a right
				newQ = append(newQ, *node.Right)
			}
		}
		height++
		q = newQ
		newQ = []TreeNode{}
	}
}

func main() {
	fmt.Println()
}
