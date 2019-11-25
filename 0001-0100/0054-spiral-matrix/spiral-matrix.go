package leetcode

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	row, column := len(matrix), len(matrix[0])

	next := nextFunc(row, column)

	result := make([]int, row*column)
	for i := range result {
		x, y := next()
		result[i] = matrix[x][y]
	}

	return result
}

func nextFunc(row, column int) func() (int, int) {
	top, down := 0, row-1
	left, right := 0, column-1
	x, y := 0, -1
	dx, dy := 0, 1
	return func() (int, int) {
		x += dx
		y += dy
		switch {
		case y+dy > right:
			top++
			dx, dy = 1, 0
		case x+dx > down:
			right--
			dx, dy = 0, -1
		case y+dy < left:
			down--
			dx, dy = -1, 0
		case x+dx < top:
			left++
			dx, dy = 0, 1
		}
		return x, y
	}
}
