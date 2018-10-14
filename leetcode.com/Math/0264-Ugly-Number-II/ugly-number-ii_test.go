package uglynumberii

import (
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func nthUglyNumber(n int) int {
	var uglyNumbers = []int{1}
	var primesBase = [3]int{2, 3, 5}
	var primesNums = [3]int{2, 3, 5}
	var primesPows = [3]int{0, 0, 0}

	// get minimum of
	min := func(n [3]int) int {
		var minimum = n[0]
		if n[1] < minimum {
			minimum = n[1]
		}
		if n[2] < minimum {
			minimum = n[2]
		}
		return minimum
	}

	for index := 1; index <= n; index++ {
		uglyNumbers = append(uglyNumbers, min(primesNums))

		for i, v := range primesNums {
			if v == uglyNumbers[index] {
				primesPows[i]++
				primesNums[i] = primesBase[i] * uglyNumbers[primesPows[i]]
			}
		}
	}
	return uglyNumbers[n-1]
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v): Expected(%+v) vs Returend(%+v)\n"
var MessageOk = "OK: Input(%+v) as Expected(%+v)\n"

var TestCases = []struct {
	input    int
	expected int
}{
	{1, 1},
	{2, 2},
	{3, 3},
	{4, 4},
	{5, 5},
	{6, 6},
	{7, 8},
	{8, 9},
	{9, 10},
	{10, 12},
	{9, 10},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := nthUglyNumber(test.input)

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
			nthUglyNumber(test.input)
		}
	}
}
