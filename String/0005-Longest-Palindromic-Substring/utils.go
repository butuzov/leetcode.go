package main

func isPalindromHelper(b []byte) bool {
	var lo, hi = 0, len(b) - 1
	for lo < hi {
		if b[lo] != b[hi] {
			return false
		}
		lo++
		hi--
	}
	return true
}

func isPalindromString(s string) bool {
	return isPalindromHelper([]byte(s))
}

func isPalindromBytes(b []byte) bool {
	return isPalindromHelper(b)
}
