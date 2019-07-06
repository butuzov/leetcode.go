package main
import (
	"strconv"
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func isPalindrome(n int) bool {

	// Simple precondition check
	switch {
	case n < 0:
		return false
	}

	var reversed int
	var rest int
	var preserved = n
	for i := 0; n > 0; i++ {
		rest = n % 10
		n = (n - rest) / 10
		reversed = reversed*10 + rest
	}

	if reversed == preserved {
		return true
	}

	return false
}

func isPalindromeString(n int) bool {

	var nString = strconv.Itoa(n)

	for i := 0; i < len(nString)/2; i++ {
		if nString[i] != nString[len(nString)-1-i] {
			return false
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
	{-1, false},
	{1, true},
	{990099, true},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := isPalindrome(test.input)

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
			isPalindrome(test.input)
		}
	}
}

func BenchmarkProblemString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			isPalindromeString(test.input)
		}
	}
}
