package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxD := 1
	leftDepth := 0
	rightDepth := 0
	if root.Left != nil {
		leftDepth = maxDepth(root.Left)
	}
	if root.Right != nil {
		rightDepth = maxDepth(root.Right)
	}
	if leftDepth > rightDepth {
		maxD += leftDepth
	} else {
		maxD += rightDepth
	}
	return maxD
}
