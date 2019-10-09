package leetcode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/suaaa7/leetcode-go/tool"
)

func TestMergeTwoLists(t *testing.T) {
	tests1 := [][]int{
		[]int{1, 2, 4},
	}
	tests2 := [][]int{
		[]int{1, 3, 4},
	}
	expecteds := [][]int{
		[]int{1, 1, 2, 3, 4, 4},
	}
	for i := 0; i < len(tests1); i++ {
		actual := tool.List2Ints(
			mergeTwoLists(tool.Ints2List(tests1[i]), tool.Ints2List(tests2[i])))
		fmt.Printf("test1 = %v test2 = %v\n", tests1[i], tests2[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
