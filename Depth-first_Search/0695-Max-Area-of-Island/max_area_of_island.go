package main

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func maxAreaOfIsland(grid [][]int) (max int) {

	var rows = len(grid)
	if rows == 0 {
		return max
	}
	var cols = len(grid[0])

	if cols == 0 {
		return max
	}

	var land func(int, int) int
	land = func(r, c int) int {

		if r >= rows || r < 0 || c >= cols || c < 0 {
			// return out of bounds
			return 0
		}

		if grid[r][c] == 0 {
			// water
			return 0
		}
		grid[r][c] = 0

		return 1 + land(r-1, c) + land(r+1, c) + land(r, c-1) + land(r, c+1)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 1 {
				if landmass := land(row, col); landmass > max {
					max = landmass
				}
			}
		}
	}

	return
}
