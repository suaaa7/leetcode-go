package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := [][]string{
		[]string{},
		[]string{},
	}
	expecteds := [][][]string{
		[][]string{},
		[][]string{},
	}
	for i := 0; i < len(tests); i++ {
		actual := groupAnagrams(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
