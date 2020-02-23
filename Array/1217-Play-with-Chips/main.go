package main

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func minCostToMoveChips(chips []int) int {
	var min [2]int

	for i := 0; i < len(chips); i++ {
		min[chips[i]%2]++
	}

	if min[0] > min[1] {
		return min[1]
	}
	return min[0]
}
