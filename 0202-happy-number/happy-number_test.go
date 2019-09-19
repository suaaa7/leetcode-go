package leetcode

import (
	"fmt"
	"testing"
)

func TestIsHappy(t *testing.T) {
	tests := []int{
		19,
	}
	expecteds := []bool{
		true,
	}
	for i := 0; i < len(tests); i++ {
		actual := isHappy(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
