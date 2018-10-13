package firstuniqchar

import (
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func firstUniqChar(s string) int {
	var str = []byte(s)

	// var indexes = make([]byte, 26)
	var occurances = make(map[byte]int, 26)

	for _, c := range str {
		if _, ok := occurances[c]; !ok {
			occurances[c] = 1
		} else {
			occurances[c]++
		}
	}

	var result = -1
	for i, c := range str {
		if occurances[c] == 1 {
			result = i
			break
		}
	}
	return result
}

/*******************************************************************************
  Test Table
*******************************************************************************/

var MessageError = "Fail: Input(%s) : Expected(%d) vs Returend(%d)"
var MessageOk = "OK: Input(%s) as Expected(%d)"

var TestCases = []struct {
	input    string
	expected int
}{
	{"", -1},
	{"leetcode", 0},
	{"loveleetcode", 2},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := firstUniqChar(test.input)

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
			firstUniqChar(test.input)
		}
	}
}
