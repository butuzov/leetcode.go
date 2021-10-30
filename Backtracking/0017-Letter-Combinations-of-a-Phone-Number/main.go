package main

import "fmt"

func letterCombinations(digits string) []string {
	var (
		words   = []string{}
		solver  func(string, []string)
		buttons = map[string]string{
			"2": "abc",
			"3": "def",
			"4": "ghi",
			"5": "jkl",
			"6": "mno",
			"7": "pqrs",
			"8": "tuv",
			"9": "wxyz",
		}
	)
	solver = func(s string, l []string) {
		if len(l) == 0 {
			if len(s) != 0 {
				words = append(words, s)
			}
			return
		}

		for i := range l[0] {
			solver(s+string(l[0][i]), l[1:])
		}
	}

	initial := []string{}
	for i := range digits {
		initial = append(initial, buttons[string(digits[i])])
	}
	solver("", initial)

	return words
}

func main() {
	fmt.Println(letterCombinations("2"))
}
