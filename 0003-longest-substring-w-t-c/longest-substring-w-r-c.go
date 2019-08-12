package leetcode

import "math"

func max(nums ...int) int {
    if len(nums) == 0 {
        panic("funciton max() requires at least one argument.")
    }
    result := nums[0]
    for i := 0; i < len(nums); i++ {
        result = int(math.Max(float64(result), float64(nums[i])))
    }
    return result
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var index [256]int
	var result, left, right int

	for left < len(s) {
		if right < len(s) && index[s[right]-'a'] == 0 {
			index[s[right]-'a']++
			right++
		} else {
			index[s[left]-'a']--
			left++
		}
		result = max(result, right-left)
	}
	return result
}
