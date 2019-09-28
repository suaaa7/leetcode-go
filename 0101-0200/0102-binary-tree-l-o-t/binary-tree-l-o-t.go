package leetcode

import (
	"github.com/suaaa7/leetcode-go/tool"
)

// TreeNode has Val, Left and Right
type TreeNode = tool.TreeNode

func levelOrderBFS(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	curNum, nextLevelNum := 1, 0
	result, tmp := [][]int{}, []int{}

	for len(queue) != 0 {
		if curNum > 0 {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
				nextLevelNum++
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
				nextLevelNum++
			}
			curNum--
			tmp = append(tmp, node.Val)
			queue = queue[1:]
		}

		if curNum == 0 {
			result = append(result, tmp)
			curNum = nextLevelNum
			nextLevelNum = 0
			tmp = []int{}
		}
	}

	return result
}

func levelOrderDFS(root *TreeNode) [][]int {
	result := [][]int{}
	dfsLevel(root, -1, &result)
	return result
}

func dfsLevel(node *TreeNode, level int, result *[][]int) {
	if node == nil {
		return
	}

	curLevel := level + 1
	for len(*result) <= curLevel {
		*result = append(*result, []int{})
	}
	(*result)[curLevel] = append((*result)[curLevel], node.Val)
	dfsLevel(node.Left, curLevel, result)
	dfsLevel(node.Right, curLevel, result)
}
