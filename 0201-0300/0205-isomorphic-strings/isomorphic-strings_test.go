package leetcode

import (
	"fmt"
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	tests1 := []string{
		"egg",
		"foo",
		"paper",
	}
	tests2 := []string{
		"add",
		"bar",
		"title",
	}
	expecteds := []bool{
		true,
		false,
		true,
	}
	for i := 0; i < len(tests1); i++ {
		actual := isIsomorphic(tests1[i], tests2[i])
		fmt.Printf("test1 = %v test2 = %v\n", tests1[i], tests2[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
