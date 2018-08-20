// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

// Abbreviate returns abbreviation of the Phraze
func Abbreviate(s string) string {

	var acronim string

	// First entry - should be a part of acronym
	checkNext := true

	for _, c := range s {

		// Creating New entry for acronym
		if checkNext {
			acronim = acronim + strings.ToUpper(string(c))
			checkNext = false
		}

		// here we marking next work to be a part of Acronym
		if c == rune(' ') || c == rune('-') {
			checkNext = true
		}
	}

	return acronim
}
