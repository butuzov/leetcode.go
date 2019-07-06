package main
import "testing"

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

// a and b is our numbers that bultiplied forms a product.
// c is a modulo of 1337.
var TestCases = []struct {
	a, b, c int
}{
	{9, 1, 9},
	{99, 91, 987},
	{993, 913, 123},
	{9999, 9901, 597},
	{99979, 99681, 677},
	{999999, 999001, 1218},
	{9998017, 9997647, 877},
	{99999999, 99990001, 475},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for i, test := range TestCases {
		actual := largestPalindromLookUpTable(i + 1)

		if actual != test.c {
			t.Errorf(MessageError, i+1, test.c, actual)
		} else {
			t.Logf(MessageOk, i+1, test.c)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := range TestCases {
			largestPalindromLookUpTable(i + 1)
		}
	}
}
