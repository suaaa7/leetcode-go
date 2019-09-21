package leetcode

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	tests := [][]int{
		[]int{1, 2, 4},
		[]int{1, 2, 4},
	}
	targets := []int{
		1,
		2,
	}
	expecteds := []int{
		0,
		1,
	}
	for i := 0; i < len(targets); i++ {
		actual := searchInsert(tests[i], targets[i])
		fmt.Printf("test = %v target = %v\n", tests[i], targets[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
