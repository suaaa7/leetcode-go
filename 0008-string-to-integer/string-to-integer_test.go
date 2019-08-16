package leetcode

import (
	"fmt"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	tests := []string{
		"0",
		"42",
		"     -42",
		"4193 with words",
		"words and 987",
		"-91283472332",
	}
	expecteds := []int{
		0,
		42,
		-42,
		4193,
		0,
		-2147483648,
	}
	for i := 0; i < len(tests); i++ {
		actual := myAtoi(tests[i])
		fmt.Printf("nums = %v\n", tests[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
