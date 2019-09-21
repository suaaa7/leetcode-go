package leetcode

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
		"",
		"wsxedc",
	}
	expecteds := []int{
		3,
		1,
		3,
		0,
		6,
	}
	for i := 0; i < len(tests); i++ {
		actual := lengthOfLongestSubstring(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
