package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	S        string
	expected string
}{
	{"abbaca", "ca"},
	{"aavvbbcad", "cad"},
	{"aababaab", "ba"},
	{"babb", "ba"},
}
