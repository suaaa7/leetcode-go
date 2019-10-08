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

// Ints2ListWithCycle convert []int to *ListNode with cycle
func Ints2ListWithCycle(nums []int, pos int) *ListNode {
	head := Ints2List(nums)
	if pos == -1 {
		return head
	}

	c := head
	for pos > 0 {
		c = c.Next
		pos--
	}

	tail := c
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = c

	return head
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
