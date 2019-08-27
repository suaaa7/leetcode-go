package leetcode

import (
	"sort"
)

func subsets(nums []int) [][]int {
	result := make([][]int, 1)
	sort.Ints(nums)
	for i := range nums {
		for _, r := range result {
			clone := make([]int, len(r), len(r)+1)
			copy(clone, r)
			clone = append(clone, nums[i])
			result = append(result, clone)
		}
	}
	return result
}
