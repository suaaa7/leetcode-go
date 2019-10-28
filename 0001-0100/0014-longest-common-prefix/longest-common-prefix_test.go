package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := [][]string{
		[]string{"flower", "flow", "flight"},
		[]string{"dog", "racecar", "car"},
	}
	expecteds := []string{
		"fl",
		"",
	}
	for i := 0; i < len(tests); i++ {
		actual := longestCommonPrefix(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
