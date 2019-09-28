package tool

// TreeNode has Val, Left and Right
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NULL is null
const NULL = -1 << 63

// Ints2Tree convert []int to *TreeNode
func Ints2Tree(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0]}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && nums[i] != NULL {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && nums[i] != NULL {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// Tree2Ints convert *TreeNode to []int
func Tree2Ints(tn *TreeNode) []int {
	result := make([]int, 0, 1024)

	queue := []*TreeNode{tn}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node == nil {
				result = append(result, NULL)
			} else {
				result = append(result, node.Val)
				queue = append(queue, node.Left, node.Right)
			}
		}
		queue = queue[size:]
	}

	i := len(result)
	for i > 0 && result[i-1] == NULL {
		i--
	}

	return result[:i]
}
