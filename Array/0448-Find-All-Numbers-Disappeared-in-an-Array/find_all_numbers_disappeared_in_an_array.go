package main

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func findDisappearedNumbers(nums []int) []int {
	var results []int

	for i := 0; i < len(nums); i++ {
		for nums[i] != nums[nums[i]-1] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i := 1; i <= len(nums); i++ {
		if nums[i-1] != i {
			results = append(results, i)
		}
	}

	return results
}
