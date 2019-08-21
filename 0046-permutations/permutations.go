package leetcode

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	var permutation []int
	var result [][]int
	used := make([]bool, len(nums))

	findPermutation(nums, 0, permutation, &result, &used)

	return result
}

func findPermutation(
	nums []int,
	index int,
	permutation []int,
	result *[][]int,
	used *[]bool,
) {
	if index == len(nums) {
		b := make([]int, len(permutation))
		copy(b, permutation)
		*result = append(*result, b)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			(*used)[i] = true
			permutation = append(permutation, nums[i])
			findPermutation(nums, index+1, permutation, result, used)
			permutation = permutation[:len(permutation)-1]
			(*used)[i] = false
		}
	}
	return
}
