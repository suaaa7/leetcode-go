package leetcode

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests1 := []string{
		"anagram",
		"rat",
		"a",
		"ac",
	}
	tests2 := []string{
		"nagaram",
		"cat",
		"b",
		"bb",
	}
	expecteds := []bool{
		true,
		false,
		false,
		false,
	}
	for i := 0; i < len(tests1); i++ {
		actual := isAnagram(tests1[i], tests2[i])
		fmt.Printf("test1 = %v test2 = %v\n", tests1[i], tests2[i])
		if actual != expecteds[i] {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
