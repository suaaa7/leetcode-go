package leetcode

import "github.com/suaaa7/leetcode-go/tool"

// ListNode has Val and Next
type ListNode = tool.ListNode

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
