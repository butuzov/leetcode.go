package main

import "bytes"

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

// Manacher’s Algorithm
func longestPalindrome(s string) string {

	var b = []byte(s)
	var t = make([]byte, (len(b)*2)+3)

	t[0], t[len(t)-2], t[len(t)-1] = '^', '-', '$'
	for i := range b {
		t[(i*2)+1] = '-'
		t[(i*2)+2] = b[i]
	}

	var n = len(t)
	var p = make([]int, n)
	var C, R int

	// maxwidth and index of maxwidth column
	var maxWidthValue, maxWidthIndex int

	for i := 1; i < n-1; i++ {
		var mirror = 2*C - i
		p[i] = 0
		if R > 1 {
			p[i] = p[mirror]
			if p[mirror] > R-i {
				p[i] = R - i
			}
		}

		for {

			var (
				iLeft  = i - 1 - p[i]
				iRight = i + 1 + p[i]
			)

			if iLeft < 0 || iRight >= n {
				break
			}

			if t[iLeft] != t[iRight] {
				break
			}

			p[i]++

			if p[i] > maxWidthValue {
				maxWidthValue, maxWidthIndex = p[i], i
			}
		}

		if i+p[i] > R {
			C, R = i, i+p[i]
		}

	}

	startAt := (maxWidthIndex - 1 - maxWidthValue) / 2
	return s[startAt : startAt+maxWidthValue]
}

// naive
func longestPalindrome_v1(s string) string {
	var b = []byte(s)
	var t = make([]byte, (len(b)*2)+1)
	var p = make([]int, (len(b)*2)+1)

	for i := range b {
		t[(i*2)+1] = b[i]
	}

	// maxwidth and index of maxwidth column
	var maxWidthvalue, maxWidthIndex int

	// populating polindrom width at this index
	var palindromeWidthAtIndex = func(index int) int {
		var maxWidth int
		for i := 1; index-i >= 0 && index+i < len(t); i++ {
			if isPalindromBytes(t[index-i : index+i+1]) {
				maxWidth = i

				if maxWidth > maxWidthvalue {
					maxWidthvalue = maxWidth
					maxWidthIndex = index
				}

				continue
			}
			break
		}
		return maxWidth
	}

	for i := 0; i < len(t); i++ {
		p[i] = palindromeWidthAtIndex(i)
	}

	var startAt = maxWidthIndex - maxWidthvalue + 1
	var result bytes.Buffer
	result.Grow(maxWidthvalue)
	for i := startAt; i < startAt+maxWidthvalue*2-1; i += 2 {
		result.WriteByte(t[i])
	}

	return result.String()
}

// naive optimized
func longestPalindrome_v2(s string) string {
	var b = []byte(s)
	var t = make([]byte, (len(b)*2)+1)
	var p = make([]int, (len(b)*2)+1)
	var placeHolder byte

	for i := range b {
		t[(i*2)+1] = b[i]
	}

	// maxwidth and index of maxwidth column
	var maxWidthvalue, maxWidthIndex int

	// populating polindrom width at this index
	var palindromeWidthAtIndex = func(index int, startAt int) int {
		var maxWidth int
		for i := 1; index-i >= 0 && index+i < len(t); i += 1 {
			if isPalindromBytes(t[index-i : index+i+1]) {
				maxWidth = i

				if maxWidth > maxWidthvalue {
					maxWidthvalue = maxWidth
					maxWidthIndex = index
				}

				continue
			}
			break
		}
		return maxWidth
	}

	for i := 0; i < len(t); i++ {
		var startAtMirrorCheck = int(1)
		if t[i] != placeHolder && i > 3 {
			startAtMirrorCheck = p[i-4]
		}
		if len(t)-i < maxWidthvalue {
			break
		}
		p[i] = palindromeWidthAtIndex(i, startAtMirrorCheck)
	}

	var startAt = maxWidthIndex - maxWidthvalue + 1
	var result bytes.Buffer
	result.Grow(maxWidthvalue)
	for i := startAt; i < startAt+maxWidthvalue*2-1; i += 2 {
		result.WriteByte(t[i])
	}

	return result.String()
}
