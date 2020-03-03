package main

import "github.com/butuzov/leetcode.go/pkg/linkedlist"

var MessageError = "Fail: Input(k(%d), list(%+v)): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(k(%d), list(%+v)) as Expected(%+v)"

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	In, Out *linkedlist.ListNode
	K       int
}{
	{
		linkedlist.MakeList([]int{1, 2, 3, 4, 5}...),
		linkedlist.MakeList([]int{4, 5, 1, 2, 3}...),
		102,
	},
	{
		linkedlist.MakeList([]int{0, 1, 2}...),
		linkedlist.MakeList([]int{2, 0, 1}...),
		4,
	},
	{
		linkedlist.MakeList([]int{0, 1, 2}...),
		linkedlist.MakeList([]int{0, 1, 2}...),
		3,
	},

	{
		linkedlist.MakeList([]int{0, 1, 2}...),
		linkedlist.MakeList([]int{0, 1, 2}...),
		0,
	},
}
