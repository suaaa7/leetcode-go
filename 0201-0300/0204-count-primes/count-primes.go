package leetcode

func countPrimes(n int) int {
	isComposite := make([]bool, n)
	for i := 2; i*i < n; i++ {
		if isComposite[i] {
			continue
		}
		for j := i * i; j < n; j = j + i {
			isComposite[j] = true
		}
	}
	var count int
	for i := 2; i < n; i++ {
		if !isComposite[i] {
			count++
		}
	}

	return count
}
