package scrabble

var scores = map[byte]int{
	'a': 1,
	'b': 3,
	'c': 3,
	'd': 2,
	'e': 1,
	'f': 4,
	'g': 2,
	'h': 4,
	'i': 1,
	'j': 8,
	'k': 5,
	'l': 1,
	'm': 3,
	'n': 1,
	'o': 1,
	'p': 3,
	'q': 10,
	'r': 1,
	's': 1,
	't': 1,
	'u': 1,
	'v': 4,
	'w': 4,
	'x': 8,
	'y': 4,
	'z': 10,
}

// Score return Scribble score of the word or sentance.
func Score(w string) int {
	var score int
	var index byte

	for i := 0; i < len(w); i++ {

		// Using ASCII we getting lowercases byte of the letter.
		switch {
		// Lowercase, nothing changed.
		case w[i] >= 97 && w[i] <= 122:
			index = w[i]
			break
		// Uppercase, convert it to lowercase
		case w[i] >= 65 && w[i] <= 90:
			index = byte(w[i] + 32)
			break
		// Hmmm... let's just panic
		default:
			panic("i don't know this char")
		}

		// Performing Simple lookup and sum of results.
		score += scores[index]
	}

	return score
}
