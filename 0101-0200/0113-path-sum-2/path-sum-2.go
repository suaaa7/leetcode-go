package leetcode

import (
	"github.com/suaaa7/leetcode-go/tool"
)

// TreeNode has Val, Left and Right
type TreeNode = tool.TreeNode

func pathSum(root *TreeNode, sum int) [][]int {
	result := [][]int{}
	path := []int{}

	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, level, sum int) {
		if node == nil {
			return
		}

		if level >= len(path) {
			path = append(path, node.Val)
		} else {
			path[level] = node.Val
		}
		sum -= node.Val

		if node.Left == nil && node.Right == nil && sum == 0 {
			tmp := make([]int, level+1)
			copy(tmp, path)
			result = append(result, tmp)
		}

		dfs(node.Left, level+1, sum)
		dfs(node.Right, level+1, sum)
	}

	dfs(root, 0, sum)

	return result
}
