package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	s        string
	k        int
	expected string
}{
	{"abcd", 2, "abcd"},
	{"deeedbbcccbdaa", 3, "aa"},
	{"pbbcggttciiippooaais", 2, "ps"},
}
