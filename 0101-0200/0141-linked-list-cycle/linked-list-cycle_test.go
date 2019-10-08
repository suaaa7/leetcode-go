package leetcode

import (
	"fmt"
	"testing"

	"github.com/suaaa7/leetcode-go/tool"
)

func TestHadCycle(t *testing.T) {
	tests1 := [][]int{
		[]int{3, 2, 0, -4},
		[]int{1},
	}
	tests2 := []int{
		1,
		-1,
	}
	expecteds := []bool{
		true,
		false,
	}
	for i := 0; i < len(tests1); i++ {
		actual := hasCycle(tool.Ints2ListWithCycle(tests1[i], tests2[i]))
		fmt.Printf("test1 = %v test2 = %v\n", tests1[i], tests2[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
