package leetcode

import (
	"fmt"
	"testing"

	"github.com/suaaa7/leetcode-go/tool"
)

func TestIsSameTree(t *testing.T) {
	tests1 := [][]int{
		[]int{1, 2, 3},
		[]int{1, 2},
		[]int{1, 2, 1},
	}
	tests2 := [][]int{
		[]int{1, 2, 3},
		[]int{1, tool.NULL, 2},
		[]int{1, 1, 2},
	}
	expecteds := []bool{
		true,
		false,
		false,
	}
	for i := 0; i < len(tests1); i++ {
		actual := isSameTree(tool.Ints2Tree(tests1[i]), tool.Ints2Tree(tests2[i]))
		fmt.Printf("test1 = %v test2 = %v\n", tests1[i], tests2[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
