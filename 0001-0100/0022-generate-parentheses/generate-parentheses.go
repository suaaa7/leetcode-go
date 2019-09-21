package leetcode

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	var result []string
	findGP(n, n, "", &result)

	return result
}

func findGP(lindex, rindex int, str string, result *[]string) {
	if lindex == 0 && rindex == 0 {
		*result = append(*result, str)
		return
	}
	if lindex > 0 {
		findGP(lindex-1, rindex, str+"(", result)
	}
	if rindex > 0 && lindex < rindex {
		findGP(lindex, rindex-1, str+")", result)
	}
}
