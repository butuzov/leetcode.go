package main

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func searchMatrix(matrix [][]int, target int) bool {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	var (
		cols = len(matrix[0])
		rows = len(matrix)
	)

	// return row, col
	var indexes = func(index int) (row int, col int) {
		row = index / cols
		col = index - (row * cols)
		return
	}

	var length = cols * rows
	var lo, hi = 0, length - 1

	var row, col int

	for lo < hi {

		mid := lo + (hi-lo)/2

		row, col = indexes(mid)

		if matrix[row][col] == target {
			return true
		}

		if matrix[row][col] > target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	row, col = indexes(lo)

	return matrix[row][col] == target
}
