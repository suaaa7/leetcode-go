package tool

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInts2List(t *testing.T) {
	tests := [][]int{
		[]int{1, 2},
		[]int{1, 2, 3},
	}
	expecteds := []*ListNode{
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		},
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}
	for i := range expecteds {
		actual := Ints2List(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}

func TestInts2ListWithCycle(t *testing.T) {
	tests := [][]int{
		[]int{1, 2},
		[]int{1, 2, 3},
	}
	expecteds := []*ListNode{
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		},
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}
	for i := range expecteds {
		actual := Ints2ListWithCycle(tests[i], -1)
		fmt.Printf("test = %v\n", tests[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}

func TestList2Ints(t *testing.T) {
	tests := []*ListNode{
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		},
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}
	expecteds := [][]int{
		[]int{1, 2},
		[]int{1, 2, 3},
	}
	for i := range expecteds {
		actual := List2Ints(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
