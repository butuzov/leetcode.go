package main
import (
	"math"
	"strconv"
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func maximumSwap(num int) int {
	var nums []int
	// byte to int conversion
	for _, i := range strconv.Itoa(num) {
		nums = append(nums, int(i-48))
	}

	if len(nums) < 2 {
		return num
	}

	var max, index int
	for i := 0; i < len(nums); i++ {
		max = nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[i] != nums[j] && nums[j] >= max {
				index = j
				max = nums[j]
			}
		}

		if index != 0 {
			nums[i], nums[index] = nums[index], nums[i]
			break
		}
	}

	max = 0
	for i := 0; i < len(nums); i++ {
		max += nums[i] * int(math.Pow10(len(nums)-i-1))
	}

	return max
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    int
	expected int
}{
	{2736, 7236},
	{9973, 9973},
	{99379, 99973},
	{1993, 9913},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := maximumSwap(test.input)

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
			maximumSwap(test.input)
		}
	}
}
