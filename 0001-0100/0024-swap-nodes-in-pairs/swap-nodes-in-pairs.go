package leetcode

import (
	"fmt"

	"github.com/suaaa7/leetcode-go/tool"
)

// ListNode has Val and Next
type ListNode = tool.ListNode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	fmt.Println(tool.List2Ints(newHead))
	newHead.Next = head
	fmt.Println(tool.List2Ints(newHead))
	return newHead
}
