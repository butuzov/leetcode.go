package main

import "github.com/butuzov/leetcode.go/pkg/binarytree"

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	p        *binarytree.TreeNode
	q        *binarytree.TreeNode
	expected bool
}{
	{
		binarytree.NewTree([]string{"1", "2l", "p", "3r"}),
		binarytree.NewTree([]string{"1", "2l", "p", "3r"}),
		true,
	},
	{
		binarytree.NewTree([]string{"1", "2l"}),
		binarytree.NewTree([]string{"1", "2r"}),
		false,
	},
	{
		binarytree.NewTree([]string{"1", "2l", "p", "1r"}),
		binarytree.NewTree([]string{"1", "1l", "p", "2r"}),
		false,
	},
}
