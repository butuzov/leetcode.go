package main
var MessageError = "Fail: A(%+v) + B(%+v) is Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: A(%+v) + B(%+v) as Expected(%+v)"

var TestCases = []struct {
	llOne    *ListNode
	llTwo    *ListNode
	llResult *ListNode
}{
	{
		makeList([]int{2, 4, 3}...),
		makeList([]int{5, 6, 4}...),
		makeList([]int{7, 0, 8}...),
	},
	{
		makeList([]int{5, 2}...),
		makeList([]int{5, 9}...),
		makeList([]int{0, 2, 1}...),
	},
	{
		makeList([]int{0}...),
		makeList([]int{0}...),
		makeList([]int{0}...),
	},
}
