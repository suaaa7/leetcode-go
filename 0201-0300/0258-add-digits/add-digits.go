package leetcode

func addDigits(num int) int {
	var result int
	for num != 0 {
		result += num % 10
		num /= 10
	}

	if result > 9 {
		result = addDigits(result)
	}

	return result
}
