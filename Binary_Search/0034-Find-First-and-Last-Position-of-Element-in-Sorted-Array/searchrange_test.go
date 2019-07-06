package main
import (
	"reflect"
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/

func searchRange(nums []int, target int) []int {
	// Defining Result
	var bounds []int = []int{-1, -1}

	// defining workng bounds for first search.
	var bound, limit int = 0, len(nums) - 1

	if len(nums) == 0 {
		return bounds
	}

	// Searching for Lower Bound of the Slice
	for bound < limit {

		index := (bound + limit) / 2

		if nums[index] == target {
			limit = index
		} else if nums[index] > target {
			limit = index
		} else if nums[index] < target {
			bound = index + 1
		}
	}

	if nums[bound] != target {
		return bounds
	}

	bounds[0] = bound

	// Reseting upper bound / limit / end of the array.
	limit = len(nums) - 1

	// Searching for Lower Bound of the Slice
	for bound < limit {

		index := (bound + limit) / 2

		// if only two numbers left.
		// y <= x < y+1
		if bound == index {
			index = limit
		}

		if nums[index] == target {
			bound = index
		} else {
			limit = index - 1
		}
	}
	bounds[1] = bound
	return bounds
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := searchRange(test.input, test.key)

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf(MessageError, test.input, test.key, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.input, test.key, test.expected)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			searchRange(test.input, test.key)
		}
	}
}
