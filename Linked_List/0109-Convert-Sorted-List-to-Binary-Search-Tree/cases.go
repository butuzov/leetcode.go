package main

import (
	"strings"

	"github.com/butuzov/leetcode.go/pkg/binarytree"
	"github.com/butuzov/leetcode.go/pkg/linkedlist"
)

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	head     *linkedlist.ListNode
	expected *binarytree.TreeNode
}{
	{
		linkedlist.MakeList(0, 1, 2, 3, 4, 5),
		binarytree.NewTree(strings.Split("3,1l,0l,p,2r,p,p,5r,4l", ",")),
	},
	{
		linkedlist.MakeList(-10, -3, 0, 5, 9),
		binarytree.NewTree(strings.Split("0,-3l,-10l,p,p,9r,5l", ",")),
	},
}
