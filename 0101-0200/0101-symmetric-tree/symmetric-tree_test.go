package leetcode

import (
	"fmt"
	"testing"

	"github.com/suaaa7/leetcode-go/tool"
)

func TestIsSameTree(t *testing.T) {
	tests := [][]int{
		[]int{1, 2, 2, 3, 4, 4, 3},
		[]int{1, 2, 2, tool.NULL, 3, tool.NULL, 3},
	}
	expecteds := []bool{
		true,
		false,
	}
	for i := 0; i < len(tests); i++ {
		actual := isSymmetric(tool.Ints2Tree(tests[i]))
		fmt.Printf("test1 = %v\n", tests[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
