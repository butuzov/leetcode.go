package nrepeatedelementinsize2narray

import (
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func repeatedNTimes(A []int) int {

	repeted := (len(A) / 2) - 1
	memory := make(map[int]int)
	for _, Value := range A {
		if v := memory[Value]; v == repeted {
			return Value
		}
		memory[Value]++
	}

	return -1
}

func repeatedNTimesArray(A []int) int {
	// not as effective as map
	repeted := (len(A) / 2)
	var memory [10001]int
	for _, Value := range A {
		memory[Value]++
		if memory[Value] == repeted {
			return Value
		}
	}
	return -1
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	expected int
	input    []int
}{
	{3, []int{1, 2, 3, 3}},
	{2, []int{2, 6, 2, 1}},
	{2, []int{2, 1, 2, 5, 3, 2}},
	{5, []int{5, 1, 5, 2, 5, 3, 5, 4}},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblemMap(t *testing.T) {
	for _, test := range TestCases {
		actual := repeatedNTimes(test.input)

		if actual != test.expected {
			t.Errorf(MessageError, test.input, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.input, test.expected)
		}
	}
}

func TestProblemArray(t *testing.T) {
	for _, test := range TestCases {
		actual := repeatedNTimesArray(test.input)

		if actual != test.expected {
			t.Errorf(MessageError, test.input, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.input, test.expected)
		}
	}
}

func BenchmarkProblemMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			repeatedNTimes(test.input)
		}
	}
}

func BenchmarkProblemArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			repeatedNTimesArray(test.input)
		}
	}
}
