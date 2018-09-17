// package main

package atbash

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("result: %s", Atbash("Testing,1 2 3, testing."))
}

// Atbash chipher
func Atbash(in string) string {

	var (
		isShifted bool
		isChar    bool
		shift     int
		iter      int
	)

	// 5 chars per block checker.
	iter = 0

	// Cipher
	buf := bytes.NewBuffer([]byte{})

	for _, i := range in {

		// Based on cipher rune we determine ANSII position of it
		// we checking next ranges 65-90 Uppercase, 97-122 Lowercase,
		// 49-58 numbers.
		switch {
		case i >= 65 && i <= 90:
			isChar, isShifted, shift = true, true, 65
		case i >= 97 && i <= 122:
			isChar, isShifted, shift = true, true, 97
		case i >= 49 && i <= 58:
			isChar, isShifted, shift = true, false, 0
		}

		if isChar {

			// If its a not number applying cipher
			if isShifted {
				// new position.
				i = rune(97 + (25 - int(i) + shift))
			}

			buf.WriteRune(i)

			// Block of 5 chars.
			iter++
			if iter%5 == 0 {
				buf.WriteString(" ")
			}

		}

		// reseting varuables
		isShifted, isChar = false, false
	}

	out := buf.String()

	// instead of using strings.Trim, let's work with slice
	if out[len(out)-1:] == " " {
		return out[:len(out)-1]
	}

	return out
}
