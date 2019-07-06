package main
import (
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func isValid(s string) bool {

	// empty string!
	if len(s) == 0 {
		return true
	}

	// hey... its definatly missmatch
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
		return isValid(string(str[1:index])) &&
			isValid(string(str[index+1:]))
	}

	return isValid(string(str[1:index]))
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    string
	expected bool
}{
	{"", true},
	{"()", true},
	{"{}", true},
	{"[]", true},
	{"[{}]", true},
	{"[{}]()", true},
	{"([{}]()", false},
	{"(([]){})", true},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := isValid(test.input)
		if actual != test.expected {
			t.Errorf(MessageError, test.input, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.input, test.expected)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			isValid(test.input)
		}
	}
}
