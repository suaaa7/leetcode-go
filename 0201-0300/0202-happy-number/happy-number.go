package leetcode

func isHappy(n int) bool {
	once, twice := n, sumOfSquares(n)
	for once != twice {
		once = sumOfSquares(once)
		twice = sumOfSquares(sumOfSquares(twice))
	}

	if once == 1 {
		return true
	}

	return false
}

func sumOfSquares(n int) int {
	var result int
	for n != 0 {
		result += (n % 10) * (n % 10)
		n /= 10
	}

	return result
}
