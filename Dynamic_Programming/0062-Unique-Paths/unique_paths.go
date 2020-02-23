package main

import "fmt"

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func uniquePaths(cols int, rows int) int {
	if rows == 0 || cols == 0 {
		return 0
	}

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

		if row == rows-1 {
			return dp(row, col+1)
		}

		if col == cols-1 {
			return dp(row+1, col)
		}

		return memo(row+1, col) + memo(row, col+1)

	}

	return memo(0, 0)
}

func main() {
	fmt.Print(uniquePaths(51, 9))
}
