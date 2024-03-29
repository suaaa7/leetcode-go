package leetcode

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	return convert(clean(str))
}

func clean(str string) (sign int, abs string) {
	str = strings.TrimSpace(str)
	if str == "" {
		return
	}

	switch str[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign, abs = 1, str
	case '+':
		sign, abs = 1, str[1:]
	case '-':
		sign, abs = -1, str[1:]
	default:
		abs = ""
		return
	}

	for i, b := range abs {
		if b < '0' || '9' < b {
			abs = abs[:i]
			break
		}
	}

	return
}

func convert(sign int, absStr string) int {
	var abs int

	for _, b := range absStr {
		abs = abs*10 + int(b-'0')
		switch {
		case sign*abs > math.MaxInt32:
			return math.MaxInt32
		case sign*abs < math.MinInt32:
			return math.MinInt32
		}
	}

	return sign * abs
}
