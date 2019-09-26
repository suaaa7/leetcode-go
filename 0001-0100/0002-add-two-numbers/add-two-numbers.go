package leetcode

import (
	"github.com/suaaa7/leetcode-go/tool"
)

// ListNode has Val and Next
type ListNode = tool.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	head := &ListNode{Val: 0, Next: nil}
	current := head
	var carry int
	for l1 != nil || l2 != nil {
		var v1, v2 int
		if l1 == nil {
			v1 = 0
		} else {
			v1 = l1.Val
		}
		if l2 == nil {
			v2 = 0
		} else {
			v2 = l2.Val
		}

		current.Next = &ListNode{Val: (v1 + v2 + carry) % 10, Next: nil}
		current = current.Next
		carry = (v1 + v2 + carry) / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry % 10, Next: nil}
	}

	return head.Next
}
