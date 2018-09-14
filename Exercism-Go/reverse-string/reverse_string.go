package reverse

import "bytes"

// String - returns revered string.
func String(in string) string {
	var out []rune

	// Step 1. We creating a slce of runes
	for _, r := range in {
		out = append(out, r)
	}

	// Step 2. Itterating in reverse order we creating buffer
	reversed := bytes.NewBuffer([]byte{})
	for i := len(out) - 1; i >= 0; i-- {
		reversed.WriteRune(out[i])
	}

	// Step 3. Return buffer as string
	return reversed.String()
}
