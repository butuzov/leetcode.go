package vowels

import (
	"bytes"
)

func reverseVowels(s string) string {
	var out []rune
	var vowels []int

	var vowelsMap = map[byte]bool{
		'a': true, 'A': true,
		'o': true, 'O': true,
		'e': true, 'E': true,
		'i': true, 'I': true,
		'u': true, 'U': true,
	}

	// Step 1. We creating a slce of runes
	reversed := bytes.NewBuffer([]byte{})
	for i, r := range s {
		out = append(out, r)
		if _, ok := vowelsMap[byte(r)]; ok {
			vowels = append(vowels, i)
		}
	}

	// Step 2. Actual Reversing
	vowelsLen := len(vowels)
	for i := vowelsLen - 1; i > vowelsLen-i-1; i-- {
		iID, oID := vowels[i], vowels[vowelsLen-i-1]
		out[iID], out[oID] = out[oID], out[iID]
	}

	// Step 3. Itterating in reverse order we creating buffer
	for i := range out {
		reversed.WriteRune(out[i])
	}

	// Step 4. Return buffer as string
	return reversed.String()
}
