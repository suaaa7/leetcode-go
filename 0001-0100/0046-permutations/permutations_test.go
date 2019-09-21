package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := [][]int{
		[]int{1, 2},
		[]int{1, 2, 3},
		[]int{},
	}
	expecteds := [][][]int{
		[][]int{[]int{1, 2}, []int{2, 1}},
		[][]int{
			[]int{1, 2, 3},
			[]int{1, 3, 2},
			[]int{2, 1, 3},
			[]int{2, 3, 1},
			[]int{3, 1, 2},
			[]int{3, 2, 1},
		},
		[][]int{},
	}
	for i := 0; i < len(tests); i++ {
		actual := permute(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
