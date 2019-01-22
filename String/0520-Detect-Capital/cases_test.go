package capitals

var MessageError = "Fail: Input(%+v) : Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    string
	expected bool
}{
	{"USA", true},
	{"FlaG", false},
	{"Google", true},
	{"leetcode", true},
}
