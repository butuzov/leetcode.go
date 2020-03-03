package main

import "github.com/butuzov/leetcode.go/pkg/linkedlist"

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v): Expected(%s) vs Returend(%s)"
var MessageOk = "OK: Input(%+v) as Expected(%s)"

var TestCases = []struct {
	input    []*linkedlist.ListNode
	expected *linkedlist.ListNode
}{
	{
		[]*linkedlist.ListNode{
			linkedlist.MakeList([]int{1}...),
			linkedlist.MakeList([]int{2}...),
		},
		linkedlist.MakeList([]int{1, 2}...),
	},
	{
		[]*linkedlist.ListNode{
			nil,
			linkedlist.MakeList([]int{2}...),
		},
		linkedlist.MakeList([]int{2}...),
	},
}
