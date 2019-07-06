package main
import (
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return len(nums)
	}

	var seen int
	for _, v := range nums {
		if v == nums[seen] {
			continue
		}
		seen++
		nums[seen] = v
	}
	return seen + 1
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v)/Output(%+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v)/Output(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    []int
	expected int
}{
	{[]int{1, 1, 2}, 2},
	{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {

		copyArray := make([]int, len(test.input))
		copy(copyArray, test.input)

		actual := removeDuplicates(copyArray)

		if actual != test.expected {
			t.Errorf(MessageError, test.input, copyArray, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.input, copyArray, test.expected)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			removeDuplicates(test.input)
		}
	}
}
