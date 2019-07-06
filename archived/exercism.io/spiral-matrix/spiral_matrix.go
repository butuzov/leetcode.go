package spiralmatrix

// SpiralMatrix generated spiral matrix n*n dimesion
func SpiralMatrix(n int) [][]int {
	var matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	var directions = []string{"right", "down", "left", "up"}
	var direction, col, row int
	var high, low = n, 1
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i

		switch directions[direction%4] {
		case "right":
			col++
			if col == (high - 1) {
				direction++
			}
		case "down":
			row++
			if row == (high - 1) {
				direction++

			}
		case "left":
			col--
			if col == (low - 1) {
				direction++
				low++
			}
		case "up":
			row--
			if row == (low - 1) {
				direction++
				high--
			}
		}
	}

	return matrix
}
