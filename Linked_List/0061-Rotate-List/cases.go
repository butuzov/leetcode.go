package rotatelist

var MessageError = "Fail: Input(k(%d), list(%+v)): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(k(%d), list(%+v)) as Expected(%+v)"

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	In, Out *ListNode
	K       int
}{
	{
		makeList([]int{1, 2, 3, 4, 5}...),
		makeList([]int{4, 5, 1, 2, 3}...),
		102,
	},
	{
		makeList([]int{0, 1, 2}...),
		makeList([]int{2, 0, 1}...),
		4,
	},
	{
		makeList([]int{0, 1, 2}...),
		makeList([]int{0, 1, 2}...),
		3,
	},

	{
		makeList([]int{0, 1, 2}...),
		makeList([]int{0, 1, 2}...),
		0,
	},
}
