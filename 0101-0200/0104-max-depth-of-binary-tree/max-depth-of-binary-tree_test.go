package leetcode

import (
	"fmt"
	"testing"

	"github.com/suaaa7/leetcode-go/tool"
)

func TestMaxDepth(t *testing.T) {
	tests := [][]int{
		[]int{3, 9, 20, tool.NULL, tool.NULL, 15, 7},
	}
	expecteds := []int{
		3,
	}
	for i := 0; i < len(tests); i++ {
		actual := maxDepth(tool.Ints2Tree(tests[i]))
		fmt.Printf("test1 = %v\n", tests[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
