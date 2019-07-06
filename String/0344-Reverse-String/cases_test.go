package main
var MessageError = "Fail: Input(%+v) : Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    string
	expected string
}{
	{"hello", "olleh"},
	{"A man, a plan, a canal: Panama", "amanaP :lanac a ,nalp a ,nam A"},
}
