package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []int{
		5,
		1,
		0,
	}
	expecteds := [][][]int{
		[][]int{
			[]int{1},
			[]int{1, 1},
			[]int{1, 2, 1},
			[]int{1, 3, 3, 1},
			[]int{1, 4, 6, 4, 1},
		},
		[][]int{[]int{1}},
		[][]int{},
	}
	for i := 0; i < len(tests); i++ {
		actual := generate(tests[i])
		fmt.Printf("test = %v\n", tests[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
