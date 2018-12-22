package containsduplicateiii

import (
	"math"
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
// clone of https://leetcode.com/problems/contains-duplicate-iii/discuss/61639

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {

	if t < 0 || k <= 0 || len(nums) < 2 {
		return false
	}

	var tf = t

	if t == 0 {
		tf++
	}

	var index int
	var buckets = make(map[int]int)
	for i, n := range nums {

		index = int(math.Floor(float64(n) / float64(tf)))

		if _, ok := buckets[index]; ok {
			return true
		}

		if v, ok := buckets[index-1]; ok && abs(n-v) <= t {
			return true
		}

		if v, ok := buckets[index+1]; ok && abs(n-v) <= t {
			return true
		}

		buckets[index] = n
		if i >= k {
			delete(buckets, nums[i-k]/tf)
		}
	}

	return false
}

func containsNearbyAlmostDuplicateFrameLookups(nums []int, k int, t int) bool {
	if t < 0 || k <= 0 || len(nums) < 2 {
		return false
	}

	var frame = make(map[int]int)
	for i, n := range nums {
		for _, value := range frame {
			if abs(n-value) <= t {
				return true
			}
		}
		frame[i] = n
		if i >= k {
			delete(frame, i-k)
		}
	}
	return false
}

func containsNearbyAlmostDuplicateNaive(nums []int, k int, t int) bool {

	if k <= 0 || len(nums) < 2 {
		return false
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+k; j++ {
			if j == len(nums) {
				break
			}

			if i-j <= k {
				diff := nums[i] - nums[j]

				if abs(diff) <= t {
					return true
				}
			}
		}
	}

	return false
}

func abs(i int) int {
	if i < 0 {
		i = 0 - i
	}
	return i
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v) for k=%d and t=%d: Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) for k=%d and t=%d as Expected(%+v)"

var TestCases = []struct {
	array    []int
	ofst     int
	abs      int
	expected bool
}{
	{[]int{7, 1, 3}, 2, 3, true},
	{[]int{2, 2}, 3, 0, true},
	{[]int{2, 1}, 1, 1, true},
	{[]int{1, 2}, 0, 1, false},
	{[]int{1, 2, 3, 1}, 3, 0, true},
	{[]int{1, 0, 1, 1}, 1, 2, true},
	{[]int{1, 5, 9, 1, 5, 9}, 2, 3, false},

	{[]int{1}, 0, 0, false},
	{[]int{1}, 1, 1, false},
	{[]int{-1, -1}, 1, 0, true},
	{[]int{-3, 3}, 2, 4, false},
	{[]int{1, 3, 8, 4, 5}, 2, 4, true},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := containsNearbyAlmostDuplicate(test.array, test.ofst, test.abs)

		if actual != test.expected {
			t.Errorf(MessageError, test.array, test.ofst, test.abs, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.array, test.ofst, test.abs, test.expected)
		}
	}
}

func TestProblemNaive(t *testing.T) {
	for _, test := range TestCases {
		actual := containsNearbyAlmostDuplicateNaive(test.array, test.ofst, test.abs)

		if actual != test.expected {
			t.Errorf(MessageError, test.array, test.ofst, test.abs, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.array, test.ofst, test.abs, test.expected)
		}
	}
}

func TestProblemFrame(t *testing.T) {
	for _, test := range TestCases {
		actual := containsNearbyAlmostDuplicateFrameLookups(test.array, test.ofst, test.abs)

		if actual != test.expected {
			t.Errorf(MessageError, test.array, test.ofst, test.abs, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.array, test.ofst, test.abs, test.expected)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			containsNearbyAlmostDuplicate(test.array, test.ofst, test.abs)
		}
	}
}

func BenchmarkProblemNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			containsNearbyAlmostDuplicateNaive(test.array, test.ofst, test.abs)
		}
	}
}

func BenchmarkProblemFrame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			containsNearbyAlmostDuplicateFrameLookups(test.array, test.ofst, test.abs)
		}
	}
}
