package leetcode

import (
	"fmt"

	"github.com/suaaa7/leetcode-go/tool"
)

// ListNode has Val and Next
type ListNode = tool.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		fmt.Printf("min l1: l1 = %v\n", tool.List2Ints(l1))
		fmt.Printf("min l1: l2 = %v\n", tool.List2Ints(l2))
		return l1
	}

	l2.Next = mergeTwoLists(l1, l2.Next)
	fmt.Printf("min l2: l1 = %v\n", tool.List2Ints(l1))
	fmt.Printf("min l2: l2 = %v\n", tool.List2Ints(l2))
	return l2
}
