package utils

import (
	"fmt"
	"strings"
)

func PrintPyramid(matrix [][]int) string {
	var rows = make([]string, len(matrix))
	for i := 0; i < len(matrix); i++ {
		rows[i] = fmt.Sprintf("%v", matrix[i])
	}

	return strings.Join(rows, "\n")
}
