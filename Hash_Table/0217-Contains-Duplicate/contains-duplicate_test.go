package containsduplicate

import (
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func containsDuplicate(nums []int) bool {
	var seen = make(map[int]bool)
	for _, i := range nums {
		if _, ok := seen[i]; !ok {
			seen[i] = true
			continue
		}
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
	array    []int
	expected bool
}{
	{[]int{1, 2, 3, 1}, true},
	{[]int{1, 2, 3, 4}, false},
	{[]int{1, 2, 3, 1}, true},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := containsDuplicate(test.array)

		if actual != test.expected {
			t.Errorf(MessageError, test.array, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.array, test.expected)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			containsDuplicate(test.array)
		}
	}
}
