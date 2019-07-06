package main
import (
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func totalHammingDistance(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	var sum int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum += hammingDistance(nums[i], nums[j])
		}
	}

	return sum
}

func hammingDistance(x int, y int) int {
	// Source - wikipedia
	var val = x ^ y
	var sum int
	for val != 0 {
		sum++
		val &= val - 1
	}
	return sum
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    []int
	expected int
}{
	{[]int{4, 14, 2}, 6},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := totalHammingDistance(test.input)

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
			totalHammingDistance(test.input)
		}
	}
}
