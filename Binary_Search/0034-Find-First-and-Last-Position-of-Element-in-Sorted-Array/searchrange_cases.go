package main
var MessageError = "Fail: Input(%+v) Key(%d) : Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) Key(%d)  as Expected(%+v)"

var TestCases = []struct {
	input    []int
	key      int
	expected []int
}{
	{
		[]int{},
		9,
		[]int{-1, -1},
	},
	{
		[]int{5, 7, 7, 8, 8, 10},
		8,
		[]int{3, 4},
	},
	{
		[]int{5, 7, 7, 8, 8, 10},
		6,
		[]int{-1, -1},
	},
}
