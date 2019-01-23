package duplicates

var MessageError = "Fail: Input(%+v) : Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    *ListNode
	expected *ListNode
}{
	{&ListNode{}, makeList(0)},
	{makeList(0, 0), makeList(0)},
	{makeList(0, 0, 1), makeList(0, 1)},
	{makeList(1, 1, 2), makeList(1, 2)},
	{makeList(1, 1, 2, 3, 3), makeList(1, 2, 3)},
}