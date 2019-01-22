package rob

var MessageError = "Fail: In(%+v) - Expected(%v) vs Returend(%v)"
var MessageOk = "OK: In(%+v) - Return as Expected(%v)"

var TestCases = []struct {
	input    []int
	expected int
}{
	{[]int{1, 2}, 2},
	{[]int{1, 2, 4, 1}, 5},
	{[]int{1, 2, 3, 1}, 4},
	{[]int{2, 7, 9, 3, 1}, 12},
	{[]int{2, 7, 9, 3, 1, 2}, 13},
	{[]int{1, 2, 3, 1, 2, 3, 1, 1, 5}, 12},
}
