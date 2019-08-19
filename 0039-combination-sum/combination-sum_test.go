package leetcode

import (
	"fmt"
	"testing"
	"reflect"
)

func TestCombinationSum(t *testing.T) {
	tests := [][]int{
		[]int{2, 3, 6, 7},
		[]int{2, 3, 5},
	}
	targets := []int{
		7,
		8,
	}
	expecteds := [][][]int{
		[][]int{[]int{2, 2, 3}, []int{7}},
		[][]int{[]int{2, 2, 2, 2}, []int{2, 3, 3}, []int{3, 5}},
	}
	for i := 0; i < len(targets); i++ {
		actual := combinationSum(tests[i], targets[i])
		fmt.Printf("test = %v target = %v\n", tests[i], targets[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
