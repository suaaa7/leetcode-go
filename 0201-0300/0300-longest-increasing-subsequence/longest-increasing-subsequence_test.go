package leetcode

import (
	"fmt"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	tests := [][]int{
		[]int{10, 9, 2, 5, 3, 7, 101, 18},
		[]int{1, 2, 3},
	}
	expecteds := []int{
		4,
		3,
	}
	for i := 0; i < len(tests); i++ {
		actual := lengthOfLIS(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
