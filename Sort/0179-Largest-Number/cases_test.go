package main
var MessageError = "Fail: Input(%+v) : Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    []int
	expected string
}{
	{[]int{500000, 5}, "5500000"},
	{[]int{500000, 1}, "5000001"},
	{[]int{500000, 8}, "8500000"},
	{[]int{5, 500000}, "5500000"},
	{[]int{1, 500000}, "5000001"},
	{[]int{8, 500000}, "8500000"},

	// Actual casess
	{[]int{0, 0}, "0"},
	{[]int{0, 0, 0}, "0"},
	{[]int{0, 1, 0}, "100"},
	{[]int{34, 3}, "343"},
	{[]int{32, 3}, "332"},
	{[]int{1, 121}, "1211"},
	{[]int{121, 1}, "1211"},
	{[]int{121, 12}, "12121"},
	{[]int{12, 121}, "12121"},
	{[]int{3, 30, 340}, "340330"},
	{[]int{3, 30, 34, 5, 9}, "9534330"},
}
