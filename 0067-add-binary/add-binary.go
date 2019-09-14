package leetcode

import (
	"strconv"
)

func addBinary(a string, b string) string {
	var longerStr string
	var shorterStrLen int
	if len(a) > len(b) {
		longerStr = a
		shorterStrLen = len(b)
	} else {
		longerStr = b
		shorterStrLen = len(a)
	}

	result := make([]int, len(longerStr))
	var ai, bi, tmp, carry int
	for i := 0; i < shorterStrLen; i++ {
		ai, _ = strconv.Atoi(string(a[len(a)-1-i]))
		bi, _ = strconv.Atoi(string(b[len(b)-1-i]))
		tmp = ai + bi + carry
		carry = 0
		if tmp > 1 {
			tmp = tmp - 2
			carry = 1
		}
		result[len(longerStr)-1-i] = tmp
	}
	for i := shorterStrLen; i < len(longerStr); i++ {
		ai, _ = strconv.Atoi(string(longerStr[len(longerStr)-1-i]))
		tmp = ai + carry
		carry = 0
		if tmp > 1 {
			tmp = tmp - 2
			carry = 1
		}
		result[len(longerStr)-1-i] = tmp
	}

	if carry == 1 {
		return "1" + makeString(result)
	}

	return makeString(result)
}

func makeString(nums []int) string {
	bytes := make([]byte, len(nums))
	for i := range bytes {
		bytes[i] = byte(nums[i]) + '0'
	}

	return string(bytes)
}
