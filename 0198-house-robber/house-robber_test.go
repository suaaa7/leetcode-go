package leetcode

import (
	"fmt"
	"testing"
)

func TestRob(t *testing.T) {
	tests := [][]int{
		[]int{1, 2, 3, 1},
		[]int{2, 7, 9, 3, 1},
		[]int{2, 1, 1, 2},
		[]int{2, 7, 9, 3, 1, 20},
	}
	expecteds := []int{
		4,
		12,
		4,
		31,
	}
	for i := 0; i < len(tests); i++ {
		actual := rob(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
