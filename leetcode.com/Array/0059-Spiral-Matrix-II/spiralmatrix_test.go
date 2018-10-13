package spiralmatrixii

import (
	"reflect"
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func generateMatrix(n int) [][]int {
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

/*******************************************************************************
  Tests Table
*******************************************************************************/
var MessageError = "Fail: Input(%+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v): Return as Expected(%+v)"

var TestCases = []struct {
	input    int
	expected [][]int
}{
	{
		input:    0,
		expected: [][]int{},
	},
	{
		input: 1,
		expected: [][]int{
			{1},
		},
	},
	{
		input: 2,
		expected: [][]int{
			{1, 2},
			{4, 3},
		},
	},
	{
		input: 3,
		expected: [][]int{
			{1, 2, 3},
			{8, 9, 4},
			{7, 6, 5},
		},
	},
	{
		input: 4,
		expected: [][]int{
			{1, 2, 3, 4},
			{12, 13, 14, 5},
			{11, 16, 15, 6},
			{10, 9, 8, 7},
		},
	},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := generateMatrix(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf(MessageError, test.input, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.input, test.expected)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			generateMatrix(test.input)
		}
	}
}
