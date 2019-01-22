package main

import "fmt"

func main() {

	for _, i := range subsets([]int{1, 2, 3}) {
		fmt.Printf("%#v\n", i)
	}

}

func subsets(nums []int) [][]int {

	// Function reused from earlier solution for another problem.

	var Combinations func(*[][]int, []int, []int)
	Combinations = func(stored *[][]int, initial, values []int) {
		for i := range values {
			var value []int = make([]int, len(initial)+1)
			copy(value, initial)
			value[len(value)-1] = values[i]
			*stored = append(*stored, value)
			Combinations(stored, value, values[i+1:])
		}
	}

	// Actual Work ...
	var intsCombinations = make([][]int, 0)
	if len(nums) == 0 {
		return intsCombinations
	}

	// populating intsCombinations
	Combinations(&intsCombinations, []int{}, nums)
	// adding empty set
	intsCombinations = append(intsCombinations, []int{})

	return intsCombinations
}
