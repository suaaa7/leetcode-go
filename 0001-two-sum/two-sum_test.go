package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := [][]int{
		[]int{3, 2, 4},
		[]int{3, 2, 4},
		[]int{0, 8, 7, 3, 3, 4, 2},
		[]int{0, 1},
	}
	targets := []int{
		6,
		5,
		11,
		1,
	}
	expecteds := [][]int{
		[]int{1, 2},
		[]int{0, 1},
		[]int{1, 3},
		[]int{0, 1},
	}
	for i := 0; i < len(targets); i++ {
		actual := twoSum(tests[i], targets[i])
		fmt.Printf("nums = %v target = %v\n", tests[i], targets[i])
		if actual[0] != expecteds[i][0] && actual[1] != expecteds[i][1] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
