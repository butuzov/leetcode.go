package main

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func search(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	var (
		lo = 0
		hi = len(nums) - 1
	)

	var stopper = func(stopAt int) func() bool {
		var counter int
		return func() bool {
			counter++
			return counter > stopAt
		}
	}(5)

	for lo < hi {

		if stopper() {
			return -100
		}

		var mid = (lo + hi) / 2

		switch {
		case nums[mid] == target:
			return mid
		case nums[lo] <= nums[mid]:
			if nums[mid] > target && nums[lo] <= target {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		case nums[hi] >= nums[mid]:
			if nums[mid] < target && nums[hi] >= target {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}

	}

	if nums[lo] == target {
		return lo
	}

	return -1
}
