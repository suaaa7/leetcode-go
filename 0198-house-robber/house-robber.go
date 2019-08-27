package leetcode

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func rob(nums []int) int {
	var result1, result2 int
	for i, v := range nums {
		if i%2 == 0 {
			result1 = max(result1+v, result2)
		} else {
			result2 = max(result1, result2+v)
		}
	}
	return max(result1, result2)
}
