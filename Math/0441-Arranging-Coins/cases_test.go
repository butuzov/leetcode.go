package main
var MessageError = "Fail: Input(%+v) : Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    int
	expected int
}{
	{5, 2},
	{15, 5},
	{254, 22},
	{1414, 52},
	{58172, 340},
	{721752, 1200},
	{1000000, 1413},
	{99999999, 14141},
	{2147483647, 65535},
}
