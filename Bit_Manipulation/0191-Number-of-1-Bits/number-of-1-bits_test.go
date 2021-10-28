package main

import (
	"testing"
)

/*******************************************************************************
  TestTable
*******************************************************************************/

var (
	MessageError = "Fail: Input(%+v): Expected(%+v) vs Returend(%+v)"
	MessageOk    = "OK: Input(%+v) as Expected(%+v)"
)

var TestCases = []struct {
	input    uint32
	expected int
}{
	{2, 1},
	{3, 2},
	{31, 5},
	{3100, 5},
	{123123, 10},
	{12312312, 17},
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
