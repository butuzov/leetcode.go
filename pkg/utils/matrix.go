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

// Immutable Slice of string
func CopySliceString(m []string) []string {
	var slice = make([]string, len(m))
	copy(slice, m)
	return slice
}

// Immutable Slice of string
func CopySliceInt(m []int) []int {
	var slice = make([]int, len(m))
	copy(slice, m)
	return slice
}
