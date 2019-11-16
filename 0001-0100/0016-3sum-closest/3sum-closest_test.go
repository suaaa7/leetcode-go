package leetcode

import (
	"fmt"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	tests := [][]int{
		[]int{-1, 2, 1, -4},
		[]int{-1, 2, 1, -4},
		[]int{-1, 2, -1},
	}
	targets := []int{
		1,
		2,
		1,
	}
	expecteds := []int{
		2,
		2,
		0,
	}
	for i := 0; i < len(targets); i++ {
		actual := threeSumClosest(tests[i], targets[i])
		fmt.Printf("test = %v target = %v\n", tests[i], targets[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
