package leetcode

import (
	"fmt"
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {
	tests := []int{
		1,
		16,
		218,
	}
	expecteds := []bool{
		true,
		true,
		false,
	}
	for i := 0; i < len(tests); i++ {
		actual := isPowerOfTwo(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
