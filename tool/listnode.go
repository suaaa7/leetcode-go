package tool

import "fmt"

// ListNode has Val and Next
type ListNode struct {
	Val  int
	Next *ListNode
}

// Ints2List convert []int to *ListNode
func Ints2List(nums []int) *ListNode {
	result := &ListNode{}
	tmp := result

	for _, v := range nums {
		tmp.Next = &ListNode{Val: v}
		tmp = tmp.Next
	}

	return result.Next
}

// List2Ints convert *listNode to []int
func List2Ints(head *ListNode) []int {
	var cnt int
	limit := 100

	result := []int{}
	for head != nil {
		cnt++
		if cnt > limit {
			message := fmt.Sprintf("limit: %d", limit)
			panic(message)
		}

		result = append(result, head.Val)
		head = head.Next
	}

	return result
}
