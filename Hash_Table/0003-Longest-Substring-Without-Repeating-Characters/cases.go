package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	s        string
	expected int
}{
	{"", 0},
	{"a", 1},
	{"aa", 1},
	{"aaa", 1},
	{"abba", 2},
	{"bbbbbb", 1},
	{"pwwkew", 3},
	{"tmmzuxt", 5},
	{"abcabcbb", 3},
	{"aabaab!bb", 3},
	{"abcdefga", 7},
	{"abcdefggfedcba", 7},
	{"abcdefggfedfba", 7},
}
