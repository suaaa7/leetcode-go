package leetcode

import (
	"bytes"
)

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	var result bytes.Buffer
	pace := numRows*2 - 2

	for i := 0; i < len(s); i += pace {
		result.WriteByte(s[i])
	}

	for i := 1; i <= numRows-2; i++ {
		result.WriteByte(s[i])

		for j := pace; j-i < len(s); j += pace {
			result.WriteByte(s[j-i])
			if i+j < len(s) {
				result.WriteByte(s[i+j])
			}
		}
	}

	for i := numRows - 1; i < len(s); i += pace {
		result.WriteByte(s[i])
	}

	return result.String()
}
