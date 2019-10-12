package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	tests := [][]int{
		[]int{10, 1, 2, 7, 6, 1, 5},
		[]int{2, 5, 2, 1, 2},
	}
	targets := []int{
		8,
		5,
	}
	expecteds := [][][]int{
		[][]int{[]int{1, 1, 6}, []int{1, 2, 5}, []int{1, 7}, []int{2, 6}},
		[][]int{[]int{1, 2, 2}, []int{5}},
	}
	for i := 0; i < len(targets); i++ {
		actual := combinationSum2(tests[i], targets[i])
		fmt.Printf("test = %v target = %v\n", tests[i], targets[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
