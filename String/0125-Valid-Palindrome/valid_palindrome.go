package main

import "bytes"

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func isValid(b byte) bool {

	var (
		isUpperCase = b >= 65 && b <= 90
		isDigit     = b >= 48 && b <= 57
	)

	if isDigit || isUpperCase {
		return true
	}

	return false
}

func isPalindrome(s string) bool {
	var in = bytes.ToUpper([]byte(s))

	var lo, hi = 0, len(in) - 1

	for lo < hi {
		if !isValid(in[lo]) {
			lo++
			continue
		}

		if !isValid(in[hi]) {
			hi--
			continue
		}

		if in[hi] != in[lo] {
			return false
		}
		lo++
		hi--

	}

	return true
}
