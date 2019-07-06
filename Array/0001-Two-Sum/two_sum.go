package main

/*******************************************************************************
  Problem Solution
*******************************************************************************/

func TwoSum(nums []int, target int) []int {
	var lookups = map[int]int{}
	var results = []int{0, 0}
	for i := range nums {
		if idx, exists := lookups[target-nums[i]]; exists {
			results[0] = idx
			results[1] = i
			break
		}
		lookups[nums[i]] = i
	}

	return results
}

func TwoSumNaive(nums []int, target int) []int {
	var results = []int{0, 0}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				results[0] = i
				results[1] = j
				break
			}
		}
	}

	return results
}

func TwoSumLeader(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if val, ok := m[target-v]; ok {
			s := []int{val, i}
			return s
		}

		m[v] = i
	}
	return nil
}
