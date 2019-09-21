package leetcode

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := [][]int{
		[]int{1, 2, 3, 1},
		[]int{1, 2, 3, 4},
		[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
	}
	expecteds := []bool{
		true,
		false,
		true,
	}
	for i := 0; i < len(tests); i++ {
		actual := containsDuplicate(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
