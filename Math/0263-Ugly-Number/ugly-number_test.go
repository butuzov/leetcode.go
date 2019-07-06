package main
import (
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func isUgly(num int) bool {

	var primeDivisors = [3]int{2, 3, 5}

	switch {
	case num <= 0:
		return false
	case num == 1:
		return true
	}

	var index = 0
	var divisor = primeDivisors[0]
	for {

		if num%divisor == 0 {
			if num == divisor {
				break
			}
			num /= divisor
		} else {
			index++
			if index == len(primeDivisors) {
				return false
			}
			divisor = primeDivisors[index]
		}
	}

	return true
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    int
	expected bool
}{
	{46, false},
	{49, false},
	{36, true},
	{35, false},
	{0, false},
	{1, true},
	{2, true},
	{3, true},
	{5, true},
	{4, true},
	{6, true},
	{7, false},
	{14, false},
	// {-231, false},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := isUgly(test.input)

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
			isUgly(test.input)
		}
	}
}
