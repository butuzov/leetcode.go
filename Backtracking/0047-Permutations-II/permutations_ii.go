package main

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func permuteUnique(nums []int) [][]int {
	var pemutations [][]int

	var elements = make([]int, len(nums))
	copy(elements, nums)

	var permutator func(pemutation []int, elements []int)
	permutator = func(pemutation []int, elements []int) {

		// tracking permutation.
		if len(pemutation) == len(nums) {
			pemutations = append(pemutations, pemutation)
			return
		}

		var seen = make(map[int]bool)

		for i := range elements {

			if _, ok := seen[elements[i]]; ok {
				continue
			}
			seen[elements[i]] = true

			var current = make([]int, len(pemutation))
			copy(current, pemutation)
			current = append(current, elements[i])

			var elementsNew = make([]int, 0, len(elements)-1)
			elementsNew = append(elementsNew, elements[0:i]...)
			elementsNew = append(elementsNew, elements[i+1:]...)

			permutator(current, elementsNew)
		}
	}

	permutator([]int{}, elements)

	return pemutations
}
