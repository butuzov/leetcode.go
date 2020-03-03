package main

import "github.com/butuzov/leetcode.go/pkg/linkedlist"

var MessageError = "Fail: Input(%+v) : Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

// Case defining struct of our tests.
type Case struct {
	list, expected *linkedlist.ListNode
}

var TestCases = []Case{
	{
		&linkedlist.ListNode{},
		linkedlist.MakeList(0),
	},
	{
		linkedlist.MakeList(0, 0),
		linkedlist.MakeList(0),
	},
	{
		linkedlist.MakeList(0, 0, 1),
		linkedlist.MakeList(0, 1),
	},
	{
		linkedlist.MakeList(1, 1, 2),
		linkedlist.MakeList(1, 2),
	},
	{
		linkedlist.MakeList(1, 1, 2, 3, 3),
		linkedlist.MakeList(1, 2, 3),
	},
}
