package main

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

// Naive and correct way to generate Pascal Triangle, its stupid and simple
// and additive operations are wayyyyyy faster.

func generateNaive(numRows int) [][]int {

	var pascal = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		pascal[i] = make([]int, i+1)
		pascal[i][0], pascal[i][i] = 1, 1

		if i < 2 {
			continue
		}

		for col := 1; col <= i/2; col++ {
			pascal[i][col] = pascal[i-1][col-1] + pascal[i-1][col]
			pascal[i][i-col] = pascal[i][col]
		}
	}

	return pascal

}
