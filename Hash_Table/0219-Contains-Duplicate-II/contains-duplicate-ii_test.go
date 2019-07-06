package main
import (
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func containsNearbyDuplicate(nums []int, k int) bool {
	var seen = make(map[int]int)
	for i, n := range nums {

		if j, ok := seen[n]; !ok {
			seen[n] = i
			continue
		} else if i-j <= k {
			return true
		} else {
			seen[n] = i
		}
	}
	return false
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v) for k=%d: Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) for k=%d as Expected(%+v)"

var TestCases = []struct {
	array    []int
	offset   int
	expected bool
}{
	{[]int{1, 2, 3, 1}, 3, true},
	{[]int{1, 0, 1, 1}, 1, true},
	{[]int{1, 2, 3, 1, 2, 3}, 2, false},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := containsNearbyDuplicate(test.array, test.offset)

		if actual != test.expected {
			t.Errorf(MessageError, test.array, test.offset, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.array, test.offset, test.expected)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			containsNearbyDuplicate(test.array, test.offset)
		}
	}
}
