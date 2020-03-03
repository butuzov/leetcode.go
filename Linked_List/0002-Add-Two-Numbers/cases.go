package main

import "github.com/butuzov/leetcode.go/pkg/linkedlist"

var MessageError = "Fail: A(%+v) + B(%+v) is Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: A(%+v) + B(%+v) as Expected(%+v)"

var TestCases = []struct {
	llOne    *linkedlist.ListNode
	llTwo    *linkedlist.ListNode
	llResult *linkedlist.ListNode
}{
	{
		linkedlist.MakeList([]int{2, 4, 3}...),
		linkedlist.MakeList([]int{5, 6, 4}...),
		linkedlist.MakeList([]int{7, 0, 8}...),
	},
	{
		linkedlist.MakeList([]int{5, 2}...),
		linkedlist.MakeList([]int{5, 9}...),
		linkedlist.MakeList([]int{0, 2, 1}...),
	},
	{
		linkedlist.MakeList([]int{0}...),
		linkedlist.MakeList([]int{0}...),
		linkedlist.MakeList([]int{0}...),
	},
}
