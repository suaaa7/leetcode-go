package leetcode

import "github.com/suaaa7/leetcode-go/tool"

// TreeNode has Val, Left and Right
type TreeNode = tool.TreeNode

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftHeight := depth(root.Left)
	rightHeight := depth(root.Right)
	return abs(leftHeight-rightHeight) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(depth(root.Left), depth(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}

	return x
}
