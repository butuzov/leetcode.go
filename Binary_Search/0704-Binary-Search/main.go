package main

import "fmt"

func main() {
	// dude, write some code...
	fmt.Println(search([]int{1, 2, 3, 5, 8, 10}, 3))
}

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {

		mid := (lo + hi) / 2

		switch {
		case nums[mid] < target:
			lo = mid + 1
		case nums[mid] > target:
			hi = mid - 1
		case nums[mid] == target:
			return mid
		}
	}

	return -1
}
