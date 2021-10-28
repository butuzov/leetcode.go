package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	digits   string
	expected []string
}{
	{"", []string{}},
	{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	{"2", []string{"a", "b", "c"}},
}
