package main

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func numEquivDominoPairs(bones [][]int) int {
	var result int

	var lookup = make([][]int, 9)
	for i := len(lookup); i > 0; i-- {
		lookup[i-1] = make([]int, 9)
	}

	for i := range bones {
		// actual check
		var a, b = bones[i][0], bones[i][1]
		var max, min = a, b
		if b > a {
			max, min = b, a
		}
		lookup[max-1][min-1]++
	}

	for i := len(lookup); i > 0; i-- {
		for j := len(lookup[i-1]); j > 0; j-- {
			if lookup[i-1][j-1] > 1 {
				lookup[i-1][j-1]--
				for lookup[i-1][j-1] > 0 {
					result += lookup[i-1][j-1]
					lookup[i-1][j-1]--
				}
			}
		}
	}

	return result
}
