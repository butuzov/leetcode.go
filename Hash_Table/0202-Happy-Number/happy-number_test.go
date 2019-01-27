package happynumber

import (
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func isHappy(n int) bool {

	// Simple memoization look
	var memo = func() func(n int) bool {
		var nums = make(map[int]bool)
		return func(n int) bool {
			if nums[n] == false {
				nums[n] = true
				return false
			}
			return true
		}
	}()

	// Split number and check the happines
	var helper func(n int) bool
	helper = func(n int) bool {

		if memo(n) {
			return false
		}

		var sum, part int
		for n > 0 {
			part = n % 10
			n /= 10
			sum += part * part
		}

		if sum == 1 {
			return true
		}

		return helper(sum)
	}

	return helper(n)
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
	{19, true},
	{11, false},
	{10, true},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := isHappy(test.input)

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
			isHappy(test.input)
		}
	}
}
