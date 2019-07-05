package interleavingstring

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func isInterleave(s1 string, s2 string, s3 string) bool {

	if len(s1)+len(s2) != len(s3) {
		return false
	}

	if (len(s1) == 0 && s2 != s3) || (len(s2) == 0 && s1 != s3) {
		return false
	}

	if (len(s1) == 0 && s2 == s3) || (len(s2) == 0 && s1 == s3) {
		return true
	}

	var (
		b1 = []byte(s1)
		b2 = []byte(s2)
		b3 = []byte(s3)
	)

	var startsFrom func(byte, int, int, int) bool
	startsFrom = func(check byte, a int, b int, c int) bool {
		if check != b3[c] {
			return false
		}

		switch {
		case c == len(b3)-1:
			return true

		case a == len(b1):
			return startsFrom(b2[b], a, b+1, c+1)

		case b == len(b2):
			return startsFrom(b1[a], a+1, b, c+1)
		}
		return startsFrom(b1[a], a+1, b, c+1) || startsFrom(b2[b], a, b+1, c+1)
	}

	return startsFrom(b1[0], 1, 0, 0) || startsFrom(b2[0], 0, 1, 0)
}
