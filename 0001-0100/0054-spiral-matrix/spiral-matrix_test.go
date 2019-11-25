package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := [][][]int{
		[][]int{},
		[][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}},
		[][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}},
	}
	expecteds := [][]int{
		[]int{},
		[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
	}
	for i := 0; i < len(tests); i++ {
		actual := spiralOrder(tests[i])
		fmt.Printf("test1 = %v\n", tests[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
