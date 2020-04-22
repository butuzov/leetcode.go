package utils

// Immutable matrixes for tests.
func CopyMatrix(m [][]int) [][]int {
	var matrix [][]int

	if len(m) == 0 {
		return matrix
	} else {
		matrix = make([][]int, 0, len(m))
	}

	if len(m[0]) == 0 {
		return matrix
	} else {
		for row := 0; row < len(m); row++ {
			matrix = append(matrix, make([]int, len(m[0])))
			copy(matrix[row], m[row])
		}
	}

	return matrix
}
