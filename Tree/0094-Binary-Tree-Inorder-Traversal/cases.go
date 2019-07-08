package main

import "strings"

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	root     *TreeNode
	expected []int
}{
	{
		newTree(strings.Split("25|12l|10l|4l|1l|p|p|p|5r|p|p|4r|0r|31r|100r", "|")),
		[]int{1, 4, 10, 12, 5, 25, 4, 0, 31, 100},
	},
	{
		newTree(strings.Split("25|12l|6l|p|17r|p|p|37r|29l|p|50r", "|")),
		[]int{6, 12, 17, 25, 29, 37, 50},
	},
	{
		newTree(strings.Split("1|2r|3l", "|")),
		[]int{1, 3, 2},
	},
}
