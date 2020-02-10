package main

import (
	"fmt"
	"strconv"
)

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func compress(chars []byte) int {

	if len(chars) <= 1 {
		return len(chars)
	}

	var index int
	var count = 1
	for i := 1; i < len(chars); i++ {

		if chars[index] != chars[i] {

			if count != 1 {
				v := []byte(strconv.Itoa(count))
				for n := range v {
					chars[index+1+n] = v[n]
				}
				index = index + 1 + len(v)
				count = 0
			} else {
				index++
				count--
			}

			chars[index] = chars[i]
		}
		count++
	}

	if count != 1 {
		v := []byte(strconv.Itoa(count))
		for n := range v {
			chars[index+1+n] = v[n]
		}
		index = index + 1 + len(v)
	} else {
		index++
	}

	chars = chars[0:index]
	return len(chars)
}

func main() {
	// var b = []byte("aabbccc")
	var b = []byte("a")
	fmt.Printf("%d - %c", compress(b), b)
}
