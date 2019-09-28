package leetcode

import (
	"github.com/suaaa7/leetcode-go/tool"
)

// TreeNode has Val, Left and Right
type TreeNode = tool.TreeNode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
