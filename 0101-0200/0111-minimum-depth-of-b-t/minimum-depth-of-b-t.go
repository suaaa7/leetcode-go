package leetcode

import "github.com/suaaa7/leetcode-go/tool"

// TreeNode has Val, Left and Right
type TreeNode = tool.TreeNode

func minDepth(root *TreeNode) int {
	switch {
	case root == nil:
		return 0
	case root.Left == nil:
		return minDepth(root.Right) + 1
	case root.Right == nil:
		return minDepth(root.Left) + 1
	default:
		return min(minDepth(root.Left), minDepth(root.Right)) + 1
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
