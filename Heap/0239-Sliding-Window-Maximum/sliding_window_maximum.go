package main

import "math"

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func maxSlidingWindow(nums []int, k int) []int {

	var n = len(nums)

	if n < k || k == 0 {
		return []int{}
	}

	var results = make([]int, 0, n-k+1)

	var max = math.MinInt32
	var idx = -1

	for i := 0; i <= n-k; i++ {

		// local search range
		if idx < i {
			max = math.MinInt32

			var startAt = i
			for startAt < i+k {
				if nums[startAt] > max {
					max, idx = nums[startAt], startAt
				}
				startAt++
			}

			results = append(results, max)
			continue
		}

		// idx of the maximum element with in new range,
		// only checking new number with in range
		var upperBound = i + k - 1
		if nums[upperBound] > max {
			max, idx = nums[upperBound], upperBound
		}
		results = append(results, max)
	}

	return results
}
