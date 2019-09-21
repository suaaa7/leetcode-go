package leetcode

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []string{
		"()",
		"(){}[]",
		"(]",
		"([)]",
		"{[]}",
		"([]",
	}
	expecteds := []bool{
		true,
		true,
		false,
		false,
		true,
		false,
	}
	for i := 0; i < len(tests); i++ {
		actual := isValid(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
