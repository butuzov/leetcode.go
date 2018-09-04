package isogram

// IsIsogram checks is string isogram
func IsIsogram(s string) bool {

	// Array index represents letter offset from starting position
	// in ascii table
	var storage [26]int
	var index int

	for _, letter := range s {

		// Using ASCII we getting lowercases byte of the letter.
		switch {
		// Lowercase, nothing changed.
		case letter >= 97 && letter <= 122:
			index = (int(letter) - 97) % 26
			break
		// Uppercase, convert it to lowercase
		case letter >= 65 && letter <= 90:
			index = (int(letter) - 65) % 26
			break
		case letter == 32 || letter == 45:
			// we ignoring space or dash
			index = -1
			break
		default:
			panic("i don't know what to do")
		}

		if index >= 0 {
			if storage[index] == 1 {
				return false
			}
			storage[index]++
		}

	}

	return true
}
