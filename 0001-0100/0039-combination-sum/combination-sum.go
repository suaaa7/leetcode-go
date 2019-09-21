package leetcode

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	var comb []int
	var result [][]int

	sort.Ints(candidates)
	findCombinationSum(candidates, target, 0, comb, &result)

	return result
}

func findCombinationSum(
	candidates []int,
	target, index int,
	comb []int,
	result *[][]int,
) {
	if target < 0 {
		return
	}
	if target == 0 {
		b := make([]int, len(comb))
		copy(b, comb)
		*result = append(*result, b)
		return
	}

	for i := index; i < len(candidates); i++ {
		comb = append(comb, candidates[i])
		findCombinationSum(candidates, target-candidates[i], i, comb, result)
		comb = comb[:len(comb)-1]
	}
}
