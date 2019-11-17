package tool

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInts2Tree(t *testing.T) {
	tests := [][]int{
		[]int{1, 2},
		[]int{1, 2, 3},
	}
	expecteds := []*TreeNode{
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		},
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	for i := range expecteds {
		actual := Ints2Tree(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}

func TestTree2Ints(t *testing.T) {
	tests := []*TreeNode{
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		},
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	expecteds := [][]int{
		[]int{1, 2},
		[]int{1, 2, 3},
	}
	for i := range expecteds {
		actual := Tree2Ints(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
