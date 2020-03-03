package main

import (
	"strings"

	"github.com/butuzov/leetcode.go/pkg/binarytree"
)

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	root     *binarytree.TreeNode
	expected bool
}{
	{binarytree.MakeNewTree([]int{25, 12, 17, 6, 37, 50, 29}...), true},
	{binarytree.MakeNewTree([]int{10, 20, 30, 40, 14, 15, 16, 19, 39, 17}...), true},
	{binarytree.NewTree(strings.Split("2|3l", "|")), false},
	{binarytree.NewTree(strings.Split("2|1r", "|")), false},
	{binarytree.NewTree(strings.Split("2|lr|p|4r|6l", "|")), false},
	{binarytree.NewTree(strings.Split("1|1r", "|")), false},
	{binarytree.NewTree(strings.Split("1|1l", "|")), false},
	{binarytree.NewTree(strings.Split("10|5l|p|15r|6l|p|20r", "|")), false},
	{binarytree.NewTree(strings.Split("10|9l|12r|p|8l", "|")), false},
	{binarytree.NewTree(strings.Split("3|1l|0l|p|2r|3r|p|p|p|5r|4l|p|6r", "|")), false},
}

// [10,5,15,null,null,6,20]
