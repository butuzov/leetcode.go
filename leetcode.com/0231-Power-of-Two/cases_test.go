package poweroftwo

var MessageError = "Fail: Input(%+v)  Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    int64
	expected bool
}{
	{0, false},
	{1, true},
	{2, true},
	{3, false},
	{8, true},
	{262144, true},
	{262143, false},
	{4294967295, false},
	{4294967296, true},
}
