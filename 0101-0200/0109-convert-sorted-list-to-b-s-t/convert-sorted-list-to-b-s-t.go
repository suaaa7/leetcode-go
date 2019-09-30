package leetcode

import (
	"github.com/suaaa7/leetcode-go/tool"
)

// ListNode has Val and Next
type ListNode = tool.ListNode

// TreeNode has Val, Left and Right
type TreeNode = tool.TreeNode

func findMiddle(head *ListNode) *ListNode {
	var prevPointer *ListNode
	slowPointer, fastPointer := head, head

	for fastPointer != nil && fastPointer.Next != nil {
		prevPointer = slowPointer
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
	}

	if prevPointer != nil {
		prevPointer.Next = nil
	}

	return slowPointer
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &TreeNode{
			Val:   head.Val,
			Left:  nil,
			Right: nil,
		}
	}

	middle := findMiddle(head)

	return &TreeNode{
		Val:   middle.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(middle.Next),
	}
}
