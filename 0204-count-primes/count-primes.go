package leetcode

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}

	if n == 3 {
		return 1
	}

	var primes []int
	primes = append(primes, 2)
	primes = append(primes, 3)

	for i := 3; i < n; i += 2 {
		if i%3 == 0 {
			continue
		}

		isPrime := true
		for j := 0; isPrime && j < len(primes); j++ {
			if i%primes[j] == 0 {
				isPrime = false
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	return len(primes)
}
