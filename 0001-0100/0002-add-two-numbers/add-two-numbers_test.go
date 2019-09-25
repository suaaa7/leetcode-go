package leetcode

import (
	"fmt"
	"testing"
)

func makeListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	result := &ListNode{Val: nums[0]}
	tmp := result

	for i := 1; i < len(nums); i++ {
		tmp.Next = &ListNode{Val: nums[i]}
		tmp = tmp.Next
	}

	return result
}

func isEqualListNode(l1 *ListNode, l2 *ListNode) bool {
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

		if v1 != v2 {
			return false
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return true
}

func TestAddTwoNumbers(t *testing.T) {
	tests1 := [][]int{
		[]int{2, 4, 3},
		[]int{5, 5, 2},
	}
	tests2 := [][]int{
		[]int{5, 6, 4},
		[]int{5, 5, 2},
	}
	expecteds := [][]int{
		[]int{7, 0, 8},
		[]int{0, 1, 5},
	}
	for i := 0; i < len(tests1); i++ {
		actual := addTwoNumbers(makeListNode(tests1[i]), makeListNode(tests2[i]))
		fmt.Printf("test1 = %v test2 = %v\n", tests1[i], tests2[i])
		expected := makeListNode(expecteds[i])
		if !isEqualListNode(actual, expected) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expected)
		}
	}
}
