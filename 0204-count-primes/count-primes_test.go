package leetcode

import (
	"fmt"
	"testing"
)

func TestCountPrimes(t *testing.T) {
	tests := []int{
		10,
		7,
		4,
	}
	expecteds := []int{
		4,
		3,
		2,
	}
	for i := 0; i < len(tests); i++ {
		actual := countPrimes(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
