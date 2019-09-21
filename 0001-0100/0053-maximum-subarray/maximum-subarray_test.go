package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	tests := [][]int{
		[]int{-2, 3, 6, -7},
		[]int{2, -3, 5},
		[]int{-2, -1},
	}
	expecteds := []int{
		9,
		5,
		-1,
	}
	for i := 0; i < len(tests); i++ {
		actual := maxSubArray(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
