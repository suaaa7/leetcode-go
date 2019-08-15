package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []string{
		"PAYPALISHIRING",
		"PAYPALISHIRING",
		"PAYPALISHIRING",
	}
	numRows := []int{
		3,
		1,
		4,
	}
	expecteds := []string{
		"PAHNAPLSIIGYIR",
		"PAYPALISHIRING",
		"PINALSIGYAHRPI",
	}
	for i := 0; i < len(tests); i++ {
		actual := convert(tests[i], numRows[i])
		fmt.Printf("nums = %v target = %v\n", tests[i], numRows[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
