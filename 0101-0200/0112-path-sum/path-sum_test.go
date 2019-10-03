package leetcode

import (
	"fmt"
	"testing"

	"github.com/suaaa7/leetcode-go/tool"
)

func TestHasPathSum(t *testing.T) {
	tests1 := [][]int{
		[]int{3, 9, 20},
		[]int{5, 4, 8, 11, tool.NULL, 13, 4, 7, 2, tool.NULL, tool.NULL, tool.NULL, 1},
	}
	tests2 := []int{
		100,
		22,
	}
	expecteds := []bool{
		false,
		true,
	}
	for i := 0; i < len(tests1); i++ {
		actual := hasPathSum(tool.Ints2Tree(tests1[i]), tests2[i])
		fmt.Printf("test1 = %v test2 = %v\n", tests1[i], tests2[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
