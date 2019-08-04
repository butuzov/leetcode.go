package main

func reverse(s string) string {

	var result = []byte(s)

	var lo, hi = 0, len(result) - 1

	for lo < hi {
		result[lo], result[hi] = result[hi], result[lo]
		lo++
		hi--
	}
	return string(result)
}
