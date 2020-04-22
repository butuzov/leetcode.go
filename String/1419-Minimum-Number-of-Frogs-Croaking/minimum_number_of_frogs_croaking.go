package main

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func minNumberOfFrogs(croakOfFrogs string) int {
	var tbl = []byte{'c', 'r', 'o', 'a', 'k'}
	var croak = map[byte]int{}

	if len(croakOfFrogs)%5 != 0 {
		return -1
	}

	for i := 0; i < len(tbl); i++ {
		croak[tbl[i]] = 0
	}

	var max int
	for i := 0; i < len(croakOfFrogs); i++ {
		croak[croakOfFrogs[i]]++

		if croak['c']-croak['k'] > max {
			max = croak['c'] - croak['k']
		}

		for j := 1; j < len(tbl); j++ {
			if croak[tbl[j-1]] < croak[tbl[j]] {
				return -1
			}
		}
	}

	return max
}
