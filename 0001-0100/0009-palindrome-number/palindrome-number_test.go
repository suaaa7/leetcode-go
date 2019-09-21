package leetcode

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []int{
		121,
		-121,
		10,
		123,
		1221,
	}
	expecteds := []bool{
		true,
		false,
		false,
		false,
		true,
	}
	for i := 0; i < len(tests); i++ {
		actual := isPalindrome(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
