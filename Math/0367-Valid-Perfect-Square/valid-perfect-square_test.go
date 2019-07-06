package main
import (
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func isPerfectSquare(square int) bool {

	if square < 0 {
		return false
	}

	var number = square
	var numberTotal int

	for i := 0; number != 0; i++ {
		numberTotal++
		number = number / 10
	}

	var squareFloat = float64(square)
	var root = 9.
	for i := 0; i < numberTotal/2; i++ {
		root *= 10
	}

	for {
		newRoot := 0.5 * (root + squareFloat/root)
		if root-newRoot < 0.1 {
			break
		}
		root = newRoot
	}

	return int(root)*int(root) == square
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	num int
	res bool
}{
	{4, true},
	{5, false},
	{808201, true},
	{1000000, true},
	{808200, false},
	{49000000, true},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestIsPerfectSquare(t *testing.T) {
	for _, test := range TestCases {
		actual := isPerfectSquare(test.num)

		if actual != test.res {
			t.Errorf(MessageError, test.num, test.res, actual)
		} else {
			t.Logf(MessageOk, test.num, test.res)
		}
	}
}

func BenchmarkIsPerfectSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			isPerfectSquare(test.num)
		}
	}
}
