package main

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************
func sortArrayByParity(A []int) []int {

	var lo, hi = 0, len(A) - 1
	var isEven, isOdd bool

	for lo < hi {

		isEven = A[lo]%2 == 1
		isOdd = A[hi]%2 == 0

		if isEven && isOdd {
			A[lo], A[hi] = A[hi], A[lo]
			lo++
			hi--
			continue
		}

		if !isEven {
			lo++
			continue
		}

		if !isOdd {
			hi--
			continue
		}
	}
	return A
}

func validate(A []int) bool {
	var posititves = true
	for i := range A {
		tmp := A[i]%2 == 0
		if tmp == posititves {
			continue
		}

		if tmp == false && posititves {
			posititves = tmp
			continue
		}

		return false
	}
	return true
}
