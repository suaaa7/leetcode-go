package leetcode

func maxArea(height []int) int {
	var maxArea int
	l, r := 0, len(height)-1
	for l < r {
		maxArea = max(maxArea, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxArea
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
