package main

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func numIslands(grid [][]byte) (total int) {

	var rows = len(grid)
	if rows == 0 {
		return total
	}
	var cols = len(grid[0])

	if cols == 0 {
		return total
	}

	var warming func(int, int)
	warming = func(row, col int) {

		if row >= rows || row < 0 || col >= cols || col < 0 {
			// return out of bounds
			return
		}

		if grid[row][col] == '0' {
			// water
			return
		}
		grid[row][col] = '0'

		warming(row-1, col)
		warming(row+1, col)
		warming(row, col-1)
		warming(row, col+1)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				warming(row, col)
				total++
			}
		}
	}

	return
}
