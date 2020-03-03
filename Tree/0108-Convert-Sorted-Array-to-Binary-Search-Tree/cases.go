package main

import "github.com/butuzov/leetcode.go/pkg/binarytree"

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	nums     []int
	expected *binarytree.TreeNode
}{
	{
		[]int{-10, -3, 0, 5, 9},
		binarytree.NewTree([]string{"0", "-3l", "-10l", "p", "p", "9r", "5l"}),
	},
	{
		[]int{0, 1, 2, 3, 4, 5},
		binarytree.NewTree([]string{"3", "1l", "0l", "p", "2r", "p", "p", "5r", "4l"}),
	},
}
