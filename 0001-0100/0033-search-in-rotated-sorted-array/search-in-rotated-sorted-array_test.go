package leetcode

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := [][]int{
		[]int{4, 5, 6, 7, 0, 1, 2},
		[]int{4, 5, 6, 7, 0, 1, 2},
	}
	targets := []int{
		0,
		3,
	}
	expecteds := []int{
		4,
		-1,
	}
	for i := 0; i < len(targets); i++ {
		actual := search(tests[i], targets[i])
		fmt.Printf("test = %v target = %v\n", tests[i], targets[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
