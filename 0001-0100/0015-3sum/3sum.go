package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for i := range nums {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			ts := nums[i] + nums[l] + nums[r]
			switch {
			case ts < 0:
				l++
			case ts > 0:
				r--
			default:
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l, r = next(nums, l, r)
			}
		}
	}

	return result
}

func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}

	return l, r
}
