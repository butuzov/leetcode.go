package main

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func isValidString(s string) bool {

	// empty string!
	if len(s) == 0 {
		return true
	}

	// hey... its defiantly missmatch
	if len(s)%2 == 1 {
		return false
	}

	str := []byte(s)

	// we are looking for
	var startes [3]byte = [...]byte{91, 40, 123}
	var stopers [3]byte = [...]byte{93, 41, 125}
	// [ - 91  and ] - 93
	// ( - 40  and ) - 41
	// { - 123 and } - 125

	var matchedBegin byte
	var matchedBeginIndex int

	for index, match := range startes {
		if match == str[0] {
			matchedBeginIndex, matchedBegin = index, match
			break
		}
	}

	// no befin match found fail.
	if matchedBegin == 0 {
		return false
	}

	var matchedEnd byte
	var index, skip int

	for i := 1; i < len(str); i++ {

		if str[i] == startes[matchedBeginIndex] {
			skip++
		}

		if str[i] == stopers[matchedBeginIndex] {
			if skip > 0 {
				skip--
			} else {
				index, matchedEnd = i, stopers[matchedBeginIndex]
				break
			}
		}
	}

	// no end match found, fail.
	if matchedEnd == 0 {
		return false
	}

	if index+1 < len(str) {
		return isValidString(string(str[1:index])) &&
			isValidString(string(str[index+1:]))
	}

	return isValidString(string(str[1:index]))
}

func isValid(str string) bool {
	var s = make([]byte, len(str))
	var last = -1
	var diff int

	for i := range str {

		switch str[i] {
		case '(', '{', '[':
			last++

			s[last] = str[i]

		case ')', '}', ']':

			if len(s) == 0 || last < 0 {
				return false
			}

			diff = int(str[i] - s[last])
			if diff >= 1 && diff <= 2 {
				last--
			} else {
				return false
			}

		default:
			return false
		}
	}

	return last == -1
}
