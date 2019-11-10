package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := [][]int{
		[]int{-1, 0, 1, 2, -1, -4},
		[]int{-1, 0, 1},
	}
	expecteds := [][][]int{
		[][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}},
		[][]int{[]int{-1, 0, 1}},
	}
	for i := 0; i < len(tests); i++ {
		actual := threeSum(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
