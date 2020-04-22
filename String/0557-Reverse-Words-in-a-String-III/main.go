package main

import "fmt"

func main() {
	fmt.Println(reverseWords("sss!1 "))
}

func Reverse(b []byte) []byte {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	return b
}

func reverseWords(str string) string {
	var s = []byte(str)

	var index, start int

	for {

		if index > len(s)-1 {
			break
		}

		if s[index] == ' ' {
			Reverse(s[start:index])
			start = index + 1
		}

		index++
	}

	Reverse(s[start:index])

	return string(s)
}
