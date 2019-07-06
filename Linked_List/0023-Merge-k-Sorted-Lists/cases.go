package main
/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v): Expected(%s) vs Returend(%s)"
var MessageOk = "OK: Input(%+v) as Expected(%s)"

var TestCases = []struct {
	input    []*ListNode
	expected *ListNode
}{
	{
		[]*ListNode{
			makeList([]int{1}...),
			makeList([]int{2}...),
		},
		makeList([]int{1, 2}...),
	},
	{
		[]*ListNode{
			nil,
			makeList([]int{2}...),
		},
		makeList([]int{2}...),
	},
}
