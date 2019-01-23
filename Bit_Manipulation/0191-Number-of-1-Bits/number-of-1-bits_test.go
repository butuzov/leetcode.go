package numberof1bits

import (
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func hammingWeight(num uint32) int {
	var ones int
	for num > 0 {
		ones++
		// Kernigan way to count Bits
		num = num & (num - 1)
		if num == 0 {
			break
		}
	}
	return ones
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    bool
	expected bool
}{
	{true, false},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := hammingWeight(test.input)

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
			hammingWeight(test.input)
		}
	}
}
