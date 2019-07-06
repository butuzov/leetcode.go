package main
/*******************************************************************************
  Problem Solution
*******************************************************************************/

// Return Boolen statment about is n power of 2 or not.
// using bitwise operator "and".
func isPowerOfTwo(n int64) bool {
	return n > 0 && n&(n-1) == 0
}
