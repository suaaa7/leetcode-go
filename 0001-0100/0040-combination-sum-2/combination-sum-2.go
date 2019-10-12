package leetcode

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	var comb []int
	var result [][]int

	sort.Ints(candidates)
	findCombinationSum2(candidates, target, 0, comb, &result)

	return result
}

func findCombinationSum2(
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
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		if target >= candidates[i] {
			comb = append(comb, candidates[i])
			findCombinationSum2(
				candidates,
				target-candidates[i],
				i+1,
				comb,
				result,
			)
			comb = comb[:len(comb)-1]
		}
	}
}
