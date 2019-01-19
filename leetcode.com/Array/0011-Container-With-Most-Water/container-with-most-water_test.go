package containerwithmostwater

import (
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func maxArea(heights []int) int {

	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	min := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}

	if len(heights) <= 1 {
		return 0
	}

	var area int

	for i := 0; i < len(heights); i++ {
		for j := i + 1; j < len(heights); j++ {
			area = max(area, (j-i)*min(heights[i], heights[j]))
		}
	}

	return area
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    []int
	expected int
}{
	{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	{[]int{1, 1}, 1},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := maxArea(test.input)

		if actual != test.expected {
			t.Errorf(MessageError, test.input, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.input, test.expected)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			maxArea(test.input)
		}
	}
}
