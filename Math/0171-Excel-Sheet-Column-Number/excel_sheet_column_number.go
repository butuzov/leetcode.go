package main

import "math"

/*******************************************************************************
  Problem Solution
*******************************************************************************/
const ABC_LENGTH = 26

func titleToNumber(s string) int {

	var pow float64
	var sum int
	var val int
	for i, b := range s {
		val = int(b-'A') + 1
		pow = math.Pow(ABC_LENGTH, float64(len(s)-i-1))
		sum += int(pow) * val
	}

	return sum
}
