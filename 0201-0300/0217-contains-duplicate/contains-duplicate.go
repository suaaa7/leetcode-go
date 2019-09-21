package leetcode

import (
	"sort"
)

func containsDuplicate(nums []int) bool {
	if len(nums) == 1 {
		return false
	}

	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}

	return false
}
