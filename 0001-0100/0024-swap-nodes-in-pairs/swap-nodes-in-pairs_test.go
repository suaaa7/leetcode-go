package leetcode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/suaaa7/leetcode-go/tool"
)

func TestSwapPairs(t *testing.T) {
	tests := [][]int{
		[]int{1, 2, 3, 4},
		[]int{1, 2, 3},
		[]int{1, 2, 3, 4, 5, 6},
	}
	expecteds := [][]int{
		[]int{2, 1, 4, 3},
		[]int{2, 1, 3},
		[]int{2, 1, 4, 3, 6, 5},
	}
	for i := 0; i < len(tests); i++ {
		actual := tool.List2Ints(swapPairs(tool.Ints2List(tests[i])))
		fmt.Printf("test = %v\n", tests[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
