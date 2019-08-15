package leetcode

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []int{
		0,
		123,
		-123,
		120,
		1534236469,
	}
	expecteds := []int{
		0,
		321,
		-321,
		21,
		0,
	}
	for i := 0; i < len(tests); i++ {
		actual := reverse(tests[i])
		fmt.Printf("nums = %v\n", tests[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
