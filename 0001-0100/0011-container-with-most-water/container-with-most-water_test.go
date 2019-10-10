package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := [][]int{
		[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		[]int{1, 1},
	}
	expecteds := []int{
		49,
		1,
	}
	for i := 0; i < len(tests); i++ {
		actual := maxArea(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
