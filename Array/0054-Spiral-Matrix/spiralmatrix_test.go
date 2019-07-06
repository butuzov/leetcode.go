package main
import (
	"reflect"
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/

func spiralOrder(matrix [][]int) []int {

	var spiral = make([]int, 0)
	var rows, cols int

	if len(matrix) > 0 {
		rows, cols = len(matrix), len(matrix[0])
	} else {
		return spiral
	}

	var row, col int

	var directions = []string{"right", "down", "left", "up"}
	var direction int

	var rowHigh, rowLow = rows - 1, 1
	var colHigh, colLow = cols - 1, 1

	for i := 1; i <= rows*cols; i++ {

		spiral = append(spiral, matrix[row][col])

		switch directions[direction%4] {
		case "right":
			col++

			if col >= colHigh {

				// one value slice
				if col > colHigh {
					col--
					row++
				}
				direction++
				colHigh--
			}
		case "down":
			row++
			if row >= rowHigh {
				direction++
				rowHigh--
			}
		case "left":
			col--

			if col <= (colLow - 1) {
				direction++
				rowLow++
			}
		case "up":
			row--

			if row <= rowLow-1 {
				direction++
				colLow++
			}
		}
	}
	return spiral
}

/*******************************************************************************
  Test Table
*******************************************************************************/

var MessageError = "Fail: In(%#v) - \nExpected(%v)\nReturend(%v)\n"
var MessageOk = "OK: In(%+v) - Return as Expected(%v)"

var TestCases = []struct {
	input    [][]int
	expected []int
}{
	{
		[][]int{
			[]int{1, 2, 3, 4},
			[]int{5, 6, 7, 8},
			[]int{9, 10, 11, 12},
			[]int{13, 14, 15, 16},
			[]int{17, 18, 19, 20},
		},
		[]int{1, 2, 3, 4, 8, 12, 16, 20, 19, 18, 17, 13, 9, 5, 6, 7, 11, 15, 14, 10},
	},
	{
		[][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		},
		[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
	},
	{
		[][]int{
			[]int{1, 2, 3, 4},
			[]int{5, 6, 7, 8},
			[]int{9, 10, 11, 12},
		},
		[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
	},
	{
		[][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
		},
		[]int{1, 2, 3, 6, 5, 4},
	},
	{
		[][]int{[]int{1, 2, 3}},
		[]int{1, 2, 3},
	},
	{
		[][]int{},
		[]int{},
	},
	{
		[][]int{[]int{1}},
		[]int{1},
	},
	{
		[][]int{[]int{1, 2}, []int{3, 4}},
		[]int{1, 2, 4, 3},
	},
	{
		[][]int{[]int{1}, []int{2}, []int{3}},
		[]int{1, 2, 3},
	},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := spiralOrder(test.input)
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
			spiralOrder(test.input)
		}
	}
}
