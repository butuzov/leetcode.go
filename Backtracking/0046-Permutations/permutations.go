package main

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func permute(nums []int) [][]int {
	// permuations slice length
	n := len(nums)
	factorial := int(1)
	for n > 0 {
		factorial *= n
		n--
	}
	pemutations := make([][]int, 0, factorial)

	elements := make([]int, len(nums))
	copy(elements, nums)

	var permutator func(pemutation []int, elements []int)
	permutator = func(pemutation []int, elements []int) {
		// tracking permutation.
		if len(pemutation) == len(nums) {
			pemutations = append(pemutations, pemutation)
			return
		}

		for i := range elements {
			current := make([]int, len(pemutation))
			copy(current, pemutation)
			current = append(current, elements[i])

			elementsNew := make([]int, 0, len(elements)-1)
			elementsNew = append(elementsNew, elements[0:i]...)
			elementsNew = append(elementsNew, elements[i+1:]...)

			permutator(current, elementsNew)
		}
	}

	permutator([]int{}, elements)

	return pemutations
}
