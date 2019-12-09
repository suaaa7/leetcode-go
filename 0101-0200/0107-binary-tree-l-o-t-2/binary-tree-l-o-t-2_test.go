package leetcode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/suaaa7/leetcode-go/tool"
)

func TestLevelOrderBottom(t *testing.T) {
	tests := [][]int{
		[]int{3, 9, 20, tool.NULL, tool.NULL, 15, 7},
		[]int{3, 9, 20, tool.NULL, tool.NULL, 15, 100},
	}
	expecteds := [][][]int{
		[][]int{[]int{15, 7}, []int{9, 20}, []int{3}},
		[][]int{[]int{15, 100}, []int{9, 20}, []int{3}},
	}
	for i := 0; i < len(tests); i++ {
		actual := levelOrderBottom(tool.Ints2Tree(tests[i]))
		fmt.Printf("test1 = %v\n", tests[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
