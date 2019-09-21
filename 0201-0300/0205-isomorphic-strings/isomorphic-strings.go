package leetcode

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sSlice := make([]int, 256)
	tSlice := make([]int, 256)

	for i := 0; i < len(s); i++ {
		if sSlice[int(s[i])] != tSlice[int(t[i])] {
			return false
		}

		sSlice[int(s[i])] = i + 1
		tSlice[int(t[i])] = i + 1
	}

	return true
}
