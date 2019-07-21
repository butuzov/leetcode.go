package main

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func findDuplicates(nums []int) []int {

	var abs = func(x int) int {
		if x < 0 {
			x *= -1
		}
		return x
	}

	var result = make([]int, 0, len(nums)/2)

	for _, number := range nums {
		var idx = abs(number) - 1

		if nums[idx] < 0 {
			result = append(result, abs(number))
			continue
		}
		nums[idx] *= -1
	}

	return result
}
