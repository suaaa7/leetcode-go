package leetcode

import (
	"reflect"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sSlice := make([]int, 256)
	tSlice := make([]int, 256)

	for i := 0; i < len(s); i++ {
		sSlice[int(s[i])]++
		tSlice[int(t[i])]++
	}

	if !reflect.DeepEqual(sSlice, tSlice) {
		return false
	}

	return true
}
