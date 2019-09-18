package leetcode

import (
	"sort"
)

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
		i++
	}

	return nums[len(nums)-1]
}
