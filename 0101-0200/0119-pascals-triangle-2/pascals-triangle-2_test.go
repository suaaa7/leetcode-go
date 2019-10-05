package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	tests := []int{
		3,
		0,
		4,
	}
	expecteds := [][]int{
		[]int{1, 3, 3, 1},
		[]int{1},
		[]int{1, 4, 6, 4, 1},
	}
	for i := 0; i < len(tests); i++ {
		actual := getRow(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
