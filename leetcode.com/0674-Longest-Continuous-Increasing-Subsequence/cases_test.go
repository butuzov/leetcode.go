package findLengthOfLCIS

var MessageError = "Fail: Input(%+v) : Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    []int
	expected int
}{
	{[]int{}, 0},
	{[]int{1}, 1},
	{[]int{1, 1, 2}, 2},
	{[]int{1, 2, 2, 5, 5}, 2},
	{[]int{1, 3, 5, 4, 7}, 3},
	{[]int{2, 2, 2, 2, 2}, 1},
}
