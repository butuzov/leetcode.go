package superuglynumber

import (
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func nthSuperUglyNumber(n int, primes []int) int {

	var uglyNumbers = []int{1}

	var primesBase = make([]int, len(primes))
	copy(primesBase, primes)

	var primesNums = make([]int, len(primes))
	copy(primesNums, primes)

	var primesPows = make([]int, len(primes))

	// get minimum of
	min := func(n []int) int {
		var minimum = n[0]

		for i := 1; i < len(n); i++ {
			if minimum > n[i] {
				minimum = n[i]
			}
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

var MessageError = "Fail: Input(%+v) Primes(%+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) Primes(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    int
	primes   []int
	expected int
}{
	{12, []int{2, 7, 13, 19}, 32},
	{1, []int{2, 3, 5}, 1},
	{2, []int{2, 3, 5}, 2},
	{3, []int{2, 3, 5}, 3},
	{4, []int{2, 3, 5}, 4},
	{5, []int{2, 3, 5}, 5},
	{6, []int{2, 3, 5}, 6},
	{7, []int{2, 3, 5}, 8},
	{8, []int{2, 3, 5}, 9},
	{9, []int{2, 3, 5}, 10},
	{10, []int{2, 3, 5}, 12},
	{9, []int{2, 3, 5}, 10},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := nthSuperUglyNumber(test.input, test.primes)

		if actual != test.expected {
			t.Errorf(MessageError, test.input, test.primes, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.input, test.primes, test.expected)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			nthSuperUglyNumber(test.input, test.primes)
		}
	}
}
