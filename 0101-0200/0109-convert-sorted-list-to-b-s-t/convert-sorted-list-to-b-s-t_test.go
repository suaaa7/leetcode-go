package leetcode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/suaaa7/leetcode-go/tool"
)

func TestSortedListToBST(t *testing.T) {
	tests := [][]int{
		[]int{-10, -3, 0, 5, 9},
	}
	expecteds := [][]int{
		[]int{0, -3, 9, -10, tool.NULL, 5},
	}
	for i := 0; i < len(tests); i++ {
		actual := tool.Tree2Ints(sortedListToBST(tool.Ints2List(tests[i])))
		fmt.Printf("test1 = %v\n", tests[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
