package main

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func shortestPalindrome(str string) string {
	var (
		kmpBytes = []byte(str + "#" + reverse(str))
		kmpTable = make([]int, len(kmpBytes))
	)

	// populating kmp table
	for i := 1; i < len(kmpBytes); i++ {
		var idxPrevVal = kmpTable[i-1]
		for idxPrevVal > 0 && kmpBytes[idxPrevVal] != kmpBytes[i] {
			idxPrevVal = kmpTable[idxPrevVal-1]
		}
		var value int
		if kmpBytes[idxPrevVal] == kmpBytes[i] {
			value = 1
		}
		kmpTable[i] = idxPrevVal + value
	}

	var pivotPoint = kmpTable[len(kmpBytes)-1]

	return reverse(str[pivotPoint:]) + str
}
