package leetcode

func searchInsert(nums []int, target int) int {
	for index, num := range nums {
		if target <= num {
			return index
		}
	}

	return len(nums)
}
