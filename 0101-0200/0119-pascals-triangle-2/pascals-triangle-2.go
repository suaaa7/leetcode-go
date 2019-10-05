package leetcode

func getRow(rowIndex int) []int {
	result := make([]int, 1, rowIndex+1)
	result[0] = 1
	if rowIndex == 0 {
		return result
	}

	for i := 0; i < rowIndex; i++ {
		result = append(result, 1)
		for j := len(result) - 2; j > 0; j-- {
			result[j] += result[j-1]
		}
	}

	return result
}
