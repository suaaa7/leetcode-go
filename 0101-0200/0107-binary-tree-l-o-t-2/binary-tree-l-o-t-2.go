package leetcode

import (
	"github.com/suaaa7/leetcode-go/tool"
)

// TreeNode has Val, Left and Right
type TreeNode = tool.TreeNode

func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level >= len(result) {
			result = append([][]int{[]int{}}, result...)
		}
		n := len(result)
		result[n-level-1] = append(result[n-level-1], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return result
}
