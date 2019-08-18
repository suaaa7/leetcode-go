package leetcode

import (
	"fmt"
	"testing"
	"reflect"
)

func TestGenerateParentheses(t *testing.T) {
	tests := []int{
		0,
		1,
		3,
	}
	expecteds := [][]string{
		{},
		{"()",},
		{
			"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()",
		},
	}
	for i := 0; i < len(tests); i++ {
		actual := generateParenthesis(tests[i])
		fmt.Printf("tests = %v\n", tests[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			fmt.Printf(
				"actual -> %v expected -> %v\n", 
				reflect.TypeOf(actual), 
				reflect.TypeOf(expecteds[i]),
			)
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
