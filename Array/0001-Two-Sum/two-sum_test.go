package twosum

import (
	"reflect"
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func twoSum(nums []int, target int) []int {
	var results []int

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				results = append(results, i, j)
				break
			}
		}
	}

	return results
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v) Target(%d): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) Target(%d) as Expected(%+v)"

var TestCases = []struct {
	input    []int
	target   int
	expected []int
}{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{-3, 4, 3, 90}, 0, []int{0, 2}},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := twoSum(test.input, test.target)

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf(MessageError, test.input, test.target, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.input, test.target, test.expected)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			twoSum(test.input, test.target)
		}
	}
}
