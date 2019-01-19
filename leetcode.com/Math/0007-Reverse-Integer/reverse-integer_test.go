package reverseinteger

import (
	"math"
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func reverse(x int) int {
	if x < 0 {
		return -reverse(-x)
	}

	var reversed, remainder int

	for x > 0 {
		remainder = x % 10
		reversed *= 10
		reversed += remainder
		x /= 10
	}

	if reversed > math.MaxInt32 || reversed < -math.MaxInt32-1 {
		return 0
	}

	return reversed
}

func reverseNaive(x int) int {
	mul := 1
	if x < 0 {
		mul = -1
		x *= mul
	}

	var rest int
	var rest_disabled bool
	var parts []int

	// future math.Pow repalcement.
	var tenth = 1

	// spliting to parts
	for {
		rest = x % 10
		x = (x - rest) / 10

		if x == 0 && rest == 0 {
			break
		}

		if rest == 0 && !rest_disabled {
			continue
		}
		rest_disabled = true
		parts = append(parts, rest)

		// exit condition
		if x == 0 {
			break
		}

		// we going to make single operation instead
		// of using math.Pow each time
		tenth *= 10
	}

	// puting it back
	var output int
	for i := range parts {
		output += parts[i] * tenth
		tenth /= 10
	}

	output *= mul
	if output < (-2<<30) || output > (2<<30-1) {
		output = 0
	}
	return output
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    int
	expected int
}{
	{0, 0},
	{-2147483648, 0},
	{1563847412, 0},
	{1534236469, 0},
	{901000, 109},
	{123, 321},
	{-123, -321},
	{120, 21},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := reverse(test.input)

		if actual != test.expected {
			t.Errorf(MessageError, test.input, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.input, test.expected)
		}
	}
}

func TestProblemNaive(t *testing.T) {
	for _, test := range TestCases {
		actual := reverseNaive(test.input)

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
			reverse(test.input)
		}
	}
}

func BenchmarkProblemNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			reverseNaive(test.input)
		}
	}
}
