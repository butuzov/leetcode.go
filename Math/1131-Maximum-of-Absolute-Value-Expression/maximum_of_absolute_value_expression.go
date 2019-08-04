package main

import (
	"math"
)

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func Abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}

// naive
func maxAbsValExpr(arr1 []int, arr2 []int) int {

	var max = math.MinInt32

	for i := 0; i < len(arr1); i++ {
		for j := i + 1; j < len(arr1); j++ {
			m := Abs(arr1[i]-arr1[j]) + Abs(arr2[i]-arr2[j]) + Abs(i-j)
			if m > max {
				max = m
			}
		}
	}

	return max
}
