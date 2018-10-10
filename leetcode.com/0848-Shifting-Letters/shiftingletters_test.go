package shiftingletters

import (
	"bytes"
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func shiftingLetters(S string, shifts []int) string {
	var sum int
	for _, i := range shifts {
		sum += i
	}

	var buf = bytes.NewBuffer([]byte{})
	for i := 0; i < len(S); i++ {
		switch {
		case S[i] >= 'a' && S[i] <= 'z':
			buf.WriteByte(byte(((int(S[i])+sum)-int('a'))%26 + int('a')))
		case S[i] >= 'A' && S[i] <= 'Z':
			buf.WriteByte(byte(((int(S[i])+sum)-int('A'))%26 + int('A')))
		default:
			buf.WriteByte(S[i])
		}

		sum -= shifts[i]
	}

	return buf.String()
}

/*******************************************************************************
  Test Table
*******************************************************************************/
var MessageError = "Fail: Input(%s)[Shifts(%+v)] : Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%s)[Shifts(%+v)] as Expected(%+v)"

var TestCases = []struct {
	str      string
	shifts   []int
	expected string
}{
	{"abc", []int{3, 5, 9}, "rpl"},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := shiftingLetters(test.str, test.shifts)

		if actual != test.expected {
			t.Errorf(MessageError, test.str, test.shifts, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.str, test.shifts, test.expected)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			shiftingLetters(test.str, test.shifts)
		}
	}
}
