package leetcode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/suaaa7/leetcode-go/tool"
)

func TestAddTwoNumbers(t *testing.T) {
	tests1 := [][]int{
		[]int{2, 4, 3},
		[]int{5, 5, 2},
	}
	tests2 := [][]int{
		[]int{5, 6, 4},
		[]int{5, 5, 2},
	}
	expecteds := [][]int{
		[]int{7, 0, 8},
		[]int{0, 1, 5},
	}
	for i := 0; i < len(tests1); i++ {
		actual := tool.List2Ints(
			addTwoNumbers(tool.Ints2ListNode(tests1[i]), tool.Ints2ListNode(tests2[i])))
		fmt.Printf("test1 = %v test2 = %v\n", tests1[i], tests2[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
