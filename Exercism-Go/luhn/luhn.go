package luhn

import "errors"

// StringToNumbers - Converts Ctring to slice of numbers
func StringToNumbers(n string) ([]int, error) {
	var numbers []int

	// Converting string to slice of numbers
	for _, r := range n {

		// ASCII codes for numbers from 48 to 57
		if r >= 48 && r <= 57 {

			// We transforming rune to actual ints and adding them to slice
			numbers = append(numbers, (int(r) - 48))
		} else {

			// if its not a space (32 dec is space in ANSII table) - failing fast.
			if r != 32 {
				return []int{}, errors.New("String contains non digit char")
			}
		}
	}

	return numbers, nil
}

// Valid transforms string i
func Valid(n string) bool {

	numbers, err := StringToNumbers(n)
	if err != nil {
		return false
	}

	// slice is too small...
	// failing.
	if len(numbers) <= 1 {
		return false
	}

	var summ int

	// getting info what number we should check
	// odds or evens
	rightmost := (len(numbers) - 1) % 2

	// for every digit
	for i := len(numbers) - 1; i >= 0; i-- {
		var digit = numbers[i]

		// this is a case of every second digit
		// based on previous check
		if i%2 != rightmost {
			digit = digit * 2
			if digit > 9 {
				digit -= 9
			}
		}
		summ += digit
	}

	// if sum divided by 10 this is valid Luhn number.
	return summ%10 == 0
}
