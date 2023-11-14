package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil {
		if root.Right == nil {
			return true
		}
		return false
	} else {
		if root.Right == nil {
			return false
		}
	}
	leftPart := make([]*TreeNode, 0, 1)
	rightPart := make([]*TreeNode, 0, 1)
	leftPart = append(leftPart, root.Left)
	rightPart = append(rightPart, root.Right)
	notAllNodesNillsInLeftPart := true
	notAllNodesNillsInRightPart := true
	for notAllNodesNillsInLeftPart && notAllNodesNillsInRightPart {
		newLeftPart := make([]*TreeNode, 0, 2*len(leftPart))
		newRightPart := make([]*TreeNode, 0, 2*len(rightPart))
		if len(leftPart) != len(rightPart) {
			return false
		}
		notAllNodesNillsInLeftPart = false
		notAllNodesNillsInRightPart = false
		for i := 0; i < len(leftPart); i++ {
			if leftPart[i] == nil {
				if rightPart[i] == nil {
					continue
				}
				return false
			} else {
				if rightPart[i] == nil {
					return false
				}
			}
			if leftPart[i].Val != rightPart[i].Val {
				return false
			}
			notAllNodesNillsInLeftPart = notAllNodesNillsInLeftPart || (leftPart[i].Left != nil)
			notAllNodesNillsInLeftPart = notAllNodesNillsInLeftPart || (leftPart[i].Right != nil)
			notAllNodesNillsInRightPart = notAllNodesNillsInRightPart || (rightPart[i].Right != nil)
			notAllNodesNillsInRightPart = notAllNodesNillsInRightPart || (rightPart[i].Left != nil)
			newLeftPart = append(newLeftPart, leftPart[i].Left)
			newLeftPart = append(newLeftPart, leftPart[i].Right)
			newRightPart = append(newRightPart, rightPart[i].Right)
			newRightPart = append(newRightPart, rightPart[i].Left)
		}
		leftPart = newLeftPart
		rightPart = newRightPart
	}
	return true
}
