package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	s        string
	expected string
}{
	{"the sky is blue", "blue is sky the"},
	{"  hello world!  ", "world! hello"},
	{"  1", "1"},
	{"  ", ""},
	{"1  2", "2 1"},
	{"1  2 ", "2 1"},
	{" 1  2", "2 1"},
	{"a good   example", "example good a"},
}
