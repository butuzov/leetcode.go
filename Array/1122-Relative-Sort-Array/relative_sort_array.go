package main

import "sort"

/*******************************************************************************
  Problem Solution
*******************************************************************************/

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var (
		keys   = make(map[int]int)
		sorted = make([]int, len(arr1))
		ptr    = len(arr1) - 1
	)

	for i := range arr2 {
		keys[arr2[i]] = 0
	}

	for i := len(arr1) - 1; i >= 0; i-- {
		if _, ok := keys[arr1[i]]; !ok {
			sorted[ptr] = arr1[i]
			ptr--
			continue
		}
		keys[arr1[i]]++
	}

	var c = 0
	for _, v := range arr2 {
		for i := 0; i < keys[v]; i++ {
			sorted[c] = v
			c++
		}
	}

	sort.Ints(sorted[ptr+1 : len(sorted)])

	return sorted
}
