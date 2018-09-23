package findLengthOfLCIS

import (
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func findLengthOfLCIS(nums []int) int {

	if len(nums) <= 1 {
		return len(nums)
	}

	var longest, current int = 1, 1

	for n := 1; n < len(nums); n++ {
		if nums[n] > nums[(n-1)] {
			current++
			if n == len(nums)-1 && current > longest {
				longest = current
			}
		} else {

			if current > longest {
				longest = current
			}
			current = 1
		}
	}

	return longest
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := findLengthOfLCIS(test.input)

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
			findLengthOfLCIS(test.input)
		}
	}
}
