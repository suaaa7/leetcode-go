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
		fmt.Println(reflect.TypeOf(actual))
		fmt.Println(reflect.TypeOf(actual[1]))
		fmt.Println(reflect.TypeOf(expecteds[i]))
		fmt.Println(reflect.TypeOf(expecteds[i][1]))
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
