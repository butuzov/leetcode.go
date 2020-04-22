package main

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func reverseWords(str string) string {
	var s = []byte(str)
	var rev = func(b []byte) []byte {
		for i := 0; i < len(b)/2; i++ {
			b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
		}
		return b
	}

	rev(s)

	var index, start int

	for {

		if index > len(s)-1 {
			break
		}
		if index == 0 && s[index] == ' ' {
			s = s[1:]
			continue
		}

		// ending space
		if index == len(s)-1 && s[index] == ' ' {
			s = s[:index]
			continue
		}

		if index > 0 && s[index] == s[index-1] && s[index] == ' ' {
			s = append(s[:index], s[index+1:]...)
			continue
		}

		if s[index] == ' ' {
			rev(s[start:index])
			start = index + 1
		}

		index++
	}

	for {
		if len(s) > 0 && s[len(s)-1] == ' ' {
			s = s[:len(s)-1]
			continue
		}
		break
	}

	rev(s[start:index])

	return string(s)
}
