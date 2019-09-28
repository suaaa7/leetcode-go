package leetcode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/suaaa7/leetcode-go/tool"
)

func TestLevelOrderBFS(t *testing.T) {
	tests := [][]int{
		[]int{3, 9, 20, tool.NULL, tool.NULL, 15, 7},
	}
	expecteds := [][][]int{
		[][]int{[]int{3}, []int{9, 20}, []int{15, 7}},
	}
	for i := 0; i < len(tests); i++ {
		actual := levelOrderBFS(tool.Ints2Tree(tests[i]))
		fmt.Printf("test1 = %v\n", tests[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}

func TestLevelOrderDFS(t *testing.T) {
	tests := [][]int{
		[]int{3, 9, 20, tool.NULL, tool.NULL, 15, 7},
	}
	expecteds := [][][]int{
		[][]int{[]int{3}, []int{9, 20}, []int{15, 7}},
	}
	for i := 0; i < len(tests); i++ {
		actual := levelOrderDFS(tool.Ints2Tree(tests[i]))
		fmt.Printf("test1 = %v\n", tests[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
