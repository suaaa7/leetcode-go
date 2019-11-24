package leetcode

import (
	"sort"
)

type sortRunes []rune

func (s sortRunes) Len() int {
	return len(s)
}

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func groupAnagrams(strs []string) [][]string {
	tmp := make(map[string][]string, len(strs))
	for _, str := range strs {
		runes := []rune(str)
		sort.Sort(sortRunes(runes))
		tmp[string(runes)] = append(tmp[string(runes)], str)
	}

	result := make([][]string, 0, len(tmp))
	for _, v := range tmp {
		result = append(result, v)
	}

	return result
}
