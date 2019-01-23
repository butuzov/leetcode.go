package countingbits

import (
	"reflect"
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func countBits(num int) []int {
	result := make([]int, num+1)

	var n int
	for i := 2; i <= num; i++ {
		n = i
		for n > 0 {
			result[i]++
			// Kernigan way to count Bits
			n = n & (n - 1)
			if n == 0 {
				break
			}
		}
	}
	return result
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    int
	expected []int
}{
	{2, []int{0, 1, 1}},
	{5, []int{0, 1, 1, 2, 1, 2}},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := countBits(test.input)
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
			countBits(test.input)
		}
	}
}
