package main

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) == 1 {
		return true
	}

	for row := 1; row <= len(matrix)-1; row++ {
		for col := 1; col < len(matrix[row]); col++ {
			if matrix[row][col] != matrix[row-1][col-1] {
				return false
			}
		}
	}
	return true
}
