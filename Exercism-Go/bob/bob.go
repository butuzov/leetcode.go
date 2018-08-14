// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {

	var isLowercase bool
	var isUppercase bool
	var isNumbers bool
	var isQuestion bool

	for _, r := range strings.Trim(remark, " ") {

		// fmt.Printf("%d %[1]c\n", r)

		// Checking All cases ( well alsmost all)
		if !isLowercase && r >= 97 && r <= 122 {
			isLowercase = true
		} else if !isUppercase && r >= 65 && r <= 90 {
			isUppercase = true
		} else if !isNumbers && r >= 48 && r <= 57 {
			isNumbers = true
		}

		// Questions? getting `char` isn't possible if you can get `rune`.
		// on last itteration its going to be true or false.
		if r == 63 {
			isQuestion = true
		} else {
			isQuestion = false
		}

	}

	// defining response.
	var response string
	if isQuestion {
		if !isLowercase && isUppercase {
			response = "Calm down, I know what I'm doing!"
		} else {
			response = "Sure."
		}
	} else {
		if !isLowercase && isUppercase {
			response = "Whoa, chill out!"
		} else if !isLowercase && !isUppercase && !isNumbers {
			response = "Fine. Be that way!"
		} else {
			response = "Whatever."
		}
	}
	return response
}
