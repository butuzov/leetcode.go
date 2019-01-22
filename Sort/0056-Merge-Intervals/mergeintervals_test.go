package mergeintervals

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
	expected []Interval
}{
	{
		[]Interval{{1, 3}, {8, 10}, {2, 6}, {15, 18}},
		[]Interval{{1, 6}, {8, 10}, {15, 18}},
	},
	{
		[]Interval{{1, 4}, {4, 5}},
		[]Interval{{1, 5}},
	},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := merge(test.input)

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
			merge(test.input)
		}
	}
}
