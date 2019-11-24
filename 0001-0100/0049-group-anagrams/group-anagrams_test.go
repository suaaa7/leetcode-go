package leetcode

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

type sortStringsSlice [][]string

func (s sortStringsSlice) Len() int {
	return len(s)
}

func (s sortStringsSlice) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func (s sortStringsSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func TestGroupAnagrams(t *testing.T) {
	tests := [][]string{
		[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
		[]string{},
	}
	expecteds := [][][]string{
		[][]string{
			[]string{"eat", "tea", "ate"},
			[]string{"tan", "nat"},
			[]string{"bat"},
		},
		[][]string{},
	}
	for i := 0; i < len(tests); i++ {
		actual := groupAnagrams(tests[i])
		sort.Sort(sortStringsSlice(actual))
		fmt.Printf("test = %v\n", tests[i])
		if !reflect.DeepEqual(actual, expecteds[i]) {
			t.Fatalf("actual -> %v expected -> %v\n", actual, expecteds[i])
		}
	}
}
