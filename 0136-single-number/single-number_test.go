package leetcode

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := [][]int{
		[]int{2},
		[]int{2, 2, 1},
		[]int{4, 1, 2, 1, 2},
	}
	expecteds := []int{
		2,
		1,
		4,
	}
	for i := 0; i < len(tests); i++ {
		actual := singleNumber(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
