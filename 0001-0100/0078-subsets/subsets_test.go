package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := [][]int{
		[]int{1, 2},
	}
	expecteds := [][][]int{
		[][]int{[]int{}, []int{1}, []int{2}, []int{1, 2}},
	}
	for i := 0; i < len(tests); i++ {
		actual := subsets(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		for j := range actual {
			if !reflect.DeepEqual(actual[j], expecteds[i][j]) && len(actual[i]) != 0 {
				t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
			}
		}
	}
}
