package leetcode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/suaaa7/leetcode-go/tool"
)

func TestPathSum(t *testing.T) {
	tests1 := [][]int{
		[]int{3, 9, 20},
		[]int{5, 4, 8, 11, tool.NULL, 13, 4, 7, 2, tool.NULL, tool.NULL, 5, 1},
	}
	tests2 := []int{
		100,
		22,
	}
	expecteds := [][][]int{
		[][]int{},
		[][]int{[]int{5, 4, 11, 2}, []int{5, 8, 4, 5}},
	}
	for i := 0; i < len(tests1); i++ {
		actual := pathSum(tool.Ints2Tree(tests1[i]), tests2[i])
		fmt.Printf("test1 = %v test2 = %v\n", tests1[i], tests2[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
