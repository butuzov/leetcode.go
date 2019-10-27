package main

import "fmt"

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func checkStraightLine(points [][]int) bool {

	if len(points) <= 1 {
		return false
	}

	if len(points) == 2 {
		return true
	}

	var results = make([]float64, 0, len(points)-1)

	for i := 1; i < len(points); i++ {
		if points[0][0] != points[i][0] && points[0][1] != points[i][1] {
			results = append(results, sidesprop(points[0], points[i]))
		}
	}

	for i := 1; i < len(results); i++ {
		if results[i] != results[i-1] {
			return false
		}
	}

	fmt.Println(results)

	return true
}

func sidesprop(a, b []int) float64 {

	var (
		width  = a[0] - b[0]
		height = a[1] - b[1]
	)

	if width < 0 {
		width = -width
	}

	if height < 0 {
		height = -height
	}

	return float64(width) / float64(height)
}
