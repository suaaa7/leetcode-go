package leetcode

import (
	"fmt"
	"testing"
)

func TestAddDigits(t *testing.T) {
	tests := []int{
		38,
		19,
		2,
	}
	expecteds := []int{
		2,
		1,
		2,
	}
	for i := 0; i < len(tests); i++ {
		actual := addDigits(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
