package leetcode

import (
	"fmt"
	"testing"
)

func TestAddBinary(t *testing.T) {
	tests1 := []string{
		"11",
		"1010",
		"1111",
	}
	tests2 := []string{
		"1",
		"1011",
		"1111",
	}
	expecteds := []string{
		"100",
		"10101",
		"11110",
	}
	for i := 0; i < len(tests1); i++ {
		actual := addBinary(tests1[i], tests2[i])
		fmt.Printf("test = %v target = %v\n", tests1[i], tests2[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
