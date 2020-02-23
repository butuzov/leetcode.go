package main

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func IsValidMatrix(m [][]int) bool {
	if len(m) == 0 {
		return false
	}
	return len(m[0]) != 0
}

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// TODO: Profile it

func minPathSumDP(m [][]int) int {

	if !IsValidMatrix(m) {
		return 0
	}

	var rows, cols = len(m), len(m[0])

	var dp func(int, int) int

	var tbl = make([][]int, rows)
	for row := 0; row < rows; row++ {
		tbl[row] = make([]int, cols)
		for col := 0; col < cols; col++ {
			tbl[row][col] = -1
		}
	}

	var memo = func(row, col int) int {

		if tbl[row][col] != -1 {

			return tbl[row][col]
		}

		tbl[row][col] = dp(row, col)
		return tbl[row][col]
	}

	dp = func(row, col int) int {

		var result int

		if row == rows-1 && col == cols-1 {
			return m[row][col]
		} else if row == rows-1 {
			result = memo(row, col+1)
		} else if col == cols-1 {
			result = memo(row+1, col)
		} else {
			result = MinInt(memo(row+1, col), memo(row, col+1))
		}

		return m[row][col] + result
	}

	return memo(0, 0)
}

func minPathSumFast(m [][]int) int {

	if !IsValidMatrix(m) {
		return 0
	}

	var rows, cols = len(m), len(m[0])

	for col := 1; col < cols; col++ {
		m[0][col] = m[0][col] + m[0][col-1]
	}
	for row := 1; row < rows; row++ {
		m[row][0] = m[row][0] + m[row-1][0]
	}

	for row := 1; row < rows; row++ {
		for col := 1; col < cols; col++ {
			m[row][col] = MinInt(m[row-1][col], m[row][col-1]) + m[row][col]
		}
	}
	return m[rows-1][cols-1]
}
