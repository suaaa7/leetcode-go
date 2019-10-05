package leetcode

func generate(numRows int) [][]int {
	result := [][]int{}
	if numRows == 0 {
		return result
	}

	result = append(result, []int{1})
	if numRows == 1 {
		return result
	}

	for i := 1; i < numRows; i++ {
		result = append(result, genNext(result[i-1]))
	}

	return result
}

func genNext(nums []int) []int {
	result := make([]int, 1, len(nums)+1)
	result = append(result, nums...)

	for i := 0; i < len(result)-1; i++ {
		result[i] += result[i+1]
	}

	return result
}
