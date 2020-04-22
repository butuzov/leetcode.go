package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	s        string
	expected string
}{
	{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
	{"Let's ", "s'teL "},
	{"aa!s ", "s!aa "},
}
