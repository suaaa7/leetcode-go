package leetcode

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []int{
		4,
		5,
		6,
	}
	expecteds := []int{
		5,
		8,
		13,
	}
	for i := 0; i < len(tests); i++ {
		actual := climbStairs(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
