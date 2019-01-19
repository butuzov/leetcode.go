package insertinterval

import (
	"reflect"
	"testing"
)

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    []Interval
	insert   Interval
	expected []Interval
}{
	{
		[]Interval{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
		Interval{4, 8},
		[]Interval{{1, 2}, {3, 10}, {12, 16}},
	},

	{
		[]Interval{},
		Interval{5, 7},
		[]Interval{{5, 7}},
	},
	{
		[]Interval{{5, 7}},
		Interval{5, 7},
		[]Interval{{5, 7}},
	},

	{
		[]Interval{{1, 2}},
		Interval{5, 7},
		[]Interval{{1, 2}, {5, 7}},
	},

	{
		[]Interval{{5, 7}},
		Interval{1, 2},
		[]Interval{{1, 2}, {5, 7}},
	},
	{
		[]Interval{{1, 3}, {6, 9}},
		Interval{2, 5},
		[]Interval{{1, 5}, {6, 9}},
	},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := insert(test.input, test.insert)

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
			insert(test.input, test.insert)
		}
	}
}
