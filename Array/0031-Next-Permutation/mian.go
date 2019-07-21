package main

import "fmt"

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}

	var pivotIndex = -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			pivotIndex = i - 1
			break
		}
	}

	if pivotIndex >= 0 {
		for i := len(nums) - 1; i > pivotIndex; i-- {
			if nums[i] > nums[pivotIndex] {
				nums[pivotIndex], nums[i] = nums[i], nums[pivotIndex]
				break
			}
		}
	}

	for i, j := (pivotIndex + 1), len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	var nums = []int{3, 2, 1}
	// var nums = []int{1, 2}
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
}
