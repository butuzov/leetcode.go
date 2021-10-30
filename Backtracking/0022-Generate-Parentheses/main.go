package main

import "fmt"

func main() {
	// dude, write some code...
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	var results []string

	var bt func(int, int, string)
	bt = func(opened, closed int, s string) {
		if len(s) == n*2 {
			results = append(results, s)
			return
		}

		if opened < n {
			bt(opened+1, closed, s+"(")
		}

		if closed < opened {
			bt(opened, closed+1, s+")")
		}
	}

	bt(0, 0, "")

	return results
}
