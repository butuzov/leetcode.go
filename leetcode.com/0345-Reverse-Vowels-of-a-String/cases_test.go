package vovels

var MessageError = "Fail: Input(%+v) : Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    string
	expected string
}{
	{"hello", "holle"},
	{"leetcode", "leotcede"},
}
