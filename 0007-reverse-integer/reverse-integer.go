package leetcode

import (
	"math"
)

func reverse(x int) int {
	var result int
	sign := 1

	if x < 0 {
		sign = -1
		x = -1 * x
	}

	for x > 0 {
		tmp := x % 10
		result = result * 10 + tmp
		x = x / 10
	}

	result = sign * result

	if result > math.MaxInt32 || result < math.MinInt32 {
		result = 0
	}

	return result
}
