package main
import (
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func rotate(nums []int, k int) {

	if len(nums) <= 1 {
		return
	}

	k = k % len(nums)

	var reverse = func(nums []int, i, j int) {
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i, j = i+1, j-1
		}
	}

	reverse(nums, 0, len(nums)-k-1)
	reverse(nums, len(nums)-k, len(nums)-1)
	reverse(nums, 0, len(nums)-1)
}

func rotateNaive(nums []int, k int) {
	var output []int

	if len(nums) <= 1 {
		return
	}

	k = k % len(nums)

	output = append(output, nums[len(nums)-k:]...)
	output = append(output, nums[:len(nums)-k]...)

	copy(nums, output)
}

func compare(X, Y []int) bool {
	if len(X) != len(Y) {
		return false
	}

	for i := range X {
		if X[i] != Y[i] {
			return false
		}
	}

	return true
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v) Step(%d): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) Step(%d) as Expected(%+v)"

var TestCases = []struct {
	input  []int
	step   int
	output []int
}{
	{
		[]int{1, 2, 3, 4, 5, 6, 7},
		3,
		[]int{5, 6, 7, 1, 2, 3, 4},
	},

	{
		[]int{-1},
		2,
		[]int{-1},
	},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {

		cpy := make([]int, len(test.input))
		copy(cpy, test.input)

		rotate(cpy, test.step)

		if !compare(cpy, test.output) {
			t.Errorf(MessageError, test.input, test.step, test.output, cpy)
		} else {
			t.Logf(MessageOk, test.input, test.step, test.output)
		}
	}
}

func TestProblemNaive(t *testing.T) {
	for _, test := range TestCases {

		cpy := make([]int, len(test.input))
		copy(cpy, test.input)

		rotateNaive(cpy, test.step)

		if !compare(cpy, test.output) {
			t.Errorf(MessageError, test.input, test.step, test.output, cpy)
		} else {
			t.Logf(MessageOk, test.input, test.step, test.output)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			rotate(test.input, test.step)
		}
	}
}

func BenchmarkProblemNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			rotateNaive(test.input, test.step)
		}
	}
}
