package leetcode

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}

	powerOfTwo := 1
	for {
		powerOfTwo *= 2
		if powerOfTwo > n {
			return false
		}

		if powerOfTwo == n {
			return true
		}
	}
}
