package main

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func reverseOnlyLetters(S string) string {
	var l, h = 0, len(S) - 1
	var b = []byte(S)

	var lowercase = func(index int) bool {
		return b[index]-'a' >= 0 && b[index]-'a' < 26
	}
	var uppercase = func(index int) bool {
		return b[index]-'A' >= 0 && b[index]-'A' < 26
	}

	for l < h {
		if !(lowercase(l) || uppercase(l)) {
			l++
			continue
		}
		if !(lowercase(h) || uppercase(h)) {
			h--
			continue
		}

		b[l], b[h] = b[h], b[l]
		l++
		h--
	}

	return string(b)
}
