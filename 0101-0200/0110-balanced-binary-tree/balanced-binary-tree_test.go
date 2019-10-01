package leetcode

import (
	"fmt"
	"testing"

	"github.com/suaaa7/leetcode-go/tool"
)

func TestIsBalanced(t *testing.T) {
	tests := [][]int{
		[]int{3, 9, 20, tool.NULL, tool.NULL, 15, 7},
		[]int{1, 2, 2, 3, 3, tool.NULL, tool.NULL, 4, 4},
	}
	expecteds := []bool{
		true,
		false,
	}
	for i := 0; i < len(tests); i++ {
		actual := isBalanced(tool.Ints2Tree(tests[i]))
		fmt.Printf("test1 = %v\n", tests[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
