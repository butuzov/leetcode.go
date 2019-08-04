package main

/*******************************************************************************
  Naive Problem Solution
*******************************************************************************/

func shortestPalindromeSimple(s string) string {
	var (
		length = len(s)
		eos    = length - 1
		idx    = 0
	)

	var w string
	for {

		w = reverse(s[eos-idx+1:]) + s
		m := length / 2

		var (
			a = w[0:m]
			b = w[m+(length%2):]
		)

		var valid = true
		for i := 0; i < len(a); i++ {
			if a[len(a)-i-1] != b[i] {
				valid = false
			}
		}

		if valid {
			break
		}

		length++
		idx++

		if idx >= len(s) {
			break
		}
	}

	return w
}
