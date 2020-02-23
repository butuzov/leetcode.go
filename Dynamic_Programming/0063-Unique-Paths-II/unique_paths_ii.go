package main

import "fmt"

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

func uniquePathsWithObstacles(m [][]int) int {

	if !IsValidMatrix(m) {
		return 0
	}

	var rows, cols = len(m), len(m[0])

	var canvas = make([][]int, rows)
	for row := 0; row < rows; row++ {
		canvas[row] = make([]int, cols)
		for col := 0; col < cols; col++ {
			canvas[row][col] = -1
		}
	}

	var dp func(int, int) int
	var memo = func(row, col int) int {
		if canvas[row][col] == -1 {
			canvas[row][col] = dp(row, col)
		}
		return canvas[row][col]
	}

	dp = func(row, col int) int {

		if row == rows || col == cols {
			return 1
		}

		if m[row][col] == 1 {
			return 0
		}

		if row == rows-1 {
			return dp(row, col+1)
		}

		if col == cols-1 {
			return dp(row+1, col)
		}
		if m[row][col] == 1 {
			return 0
		}

		return memo(row+1, col) + memo(row, col+1)

	}

	return memo(0, 0)
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}))
}
