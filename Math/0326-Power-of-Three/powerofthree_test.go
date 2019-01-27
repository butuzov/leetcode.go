package powerofthree

import (
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func isPowerOfThree(n int) bool {

	// 3486784401 = 3^20
	if n > 0 && 3486784401%n == 0 {
		return true
	}
	return false
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
	{1, true},
	{3, true},
	{9, true},
	{27, true},
	{81, true},
	{82, false},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := isPowerOfThree(test.input)

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
			isPowerOfThree(test.input)
		}
	}
}
