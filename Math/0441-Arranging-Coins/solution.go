package main
import "math"

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func arrangeCoins1(n int) int {
	var stairs int = 0
	var stair int = 0
	for n > 0 {
		stair++
		if n < stair {
			break
		}
		stairs++
		n -= stair
	}
	return stairs
}

// Return Boolen statment about is n power of 2 or not.
func arrangeCoins2(n int) int {
	// explanations available in readme.md
	return int((float64(-1) + math.Sqrt(float64(1+8*n))) / 2)
}
