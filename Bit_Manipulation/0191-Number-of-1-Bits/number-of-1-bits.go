package main

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func hammingWeight(num uint32) int {
	var ones int
	for num > 0 {
		ones++
		// Kernigan way to count Bits
		num &= (num - 1)
		if num == 0 {
			break
		}
	}
	return ones
}
