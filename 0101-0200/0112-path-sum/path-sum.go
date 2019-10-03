package leetcode

import "github.com/suaaa7/leetcode-go/tool"

// TreeNode has Val, Left and Right
type TreeNode = tool.TreeNode

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	sum -= root.Val

	if root.Left == nil && root.Right == nil {
		return sum == 0
	}

	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
