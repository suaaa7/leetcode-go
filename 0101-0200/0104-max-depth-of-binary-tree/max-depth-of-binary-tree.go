package leetcode

import (
	"github.com/suaaa7/leetcode-go/tool"
)

// TreeNode has Val, Left and Right
type TreeNode = tool.TreeNode

func maxDepth(root *TreeNode) int {
	return countDepth(root, 0)
}

func countDepth(t *TreeNode, depth int) int {
	if t == nil {
		return depth
	}

	depth++
	if t.Left == nil && t.Right == nil {
		return depth
	}

	return max(countDepth(t.Left, depth), countDepth(t.Right, depth))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
