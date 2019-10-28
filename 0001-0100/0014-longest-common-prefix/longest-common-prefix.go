package leetcode

func longestCommonPrefix(strs []string) string {
	short := shortest(strs)

	for i, s := range short {
		for i2, _ := range strs {
			if strs[i2][i] != byte(s) {
				return strs[i2][:i]
			}
		}
	}

	return short
}

func shortest(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	result := strs[0]
	for _, s := range strs {
		if len(result) > len(s) {
			result = s
		}
	}

	return result
}
