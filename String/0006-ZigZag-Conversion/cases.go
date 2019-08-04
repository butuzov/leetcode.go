package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	s        string
	numRows  int
	expected string
}{
	{"PAYPALISHIRING", 1, "PAYPALISHIRING"},
	{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
	{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
}
