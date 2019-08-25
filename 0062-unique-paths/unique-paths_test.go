package leetcode

import (
	"fmt"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	tests1 := []int{
		3,
		7,
	}
	tests2 := []int{
		2,
		3,
	}
	expecteds := []int{
		3,
		28,
	}
	for i := 0; i < len(tests1); i++ {
		actual := uniquePaths(tests1[i], tests2[i])
		fmt.Printf("test = %v target = %v\n", tests1[i], tests2[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
