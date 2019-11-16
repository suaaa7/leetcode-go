package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result, delta := 0, math.MaxInt64

	for i := range nums {
		if i > 0 && nums[i] == nums[i-i] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s < target:
				l++
				if delta > target-s {
					delta = target - s
					result = s
				}
			case s > target:
				r--
				if delta > s-target {
					delta = s - target
					result = s
				}
			default:
				return target
			}
		}
	}

	return result
}
