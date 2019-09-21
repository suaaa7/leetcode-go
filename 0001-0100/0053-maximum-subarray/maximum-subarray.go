package leetcode

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	sum, result := -1<<31, -1<<31

	for _, n := range nums {
		sum = max(sum+n, n)
		result = max(result, sum)
	}

	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
