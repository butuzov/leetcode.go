package strand

import "strings"

var runes map[rune]rune = map[rune]rune{
	rune('A'): rune('U'),
	rune('T'): rune('A'),
	rune('G'): rune('C'),
	rune('C'): rune('G'),
}

// ToRNA return RNA string.
// - we not checking is orriginal string is correct DNA, but we don't have such
//   requirment.
func ToRNA(dna string) string {
	var rna strings.Builder
	for _, i := range dna {
		rna.WriteRune(runes[i])
	}
	return rna.String()
}
