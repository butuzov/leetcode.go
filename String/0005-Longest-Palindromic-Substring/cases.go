package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	s        string
	expected string
}{
	{"ccc", "ccc"},
	{"cbbd", "bb"},
	{"cbbd", "bb"},
	{"caba", "aba"},
	{"babad", "aba"},
	{"bananas", "anana"},
	{"babadada", "adada"},
	{"aaabaaaa", "aaabaaa"},
	{"abacdfgdcaba", "aba"},
	{"ababababababa", "ababababababa"},
	{"abbadabbbadpwpw", "dabbbad"},
	{"abbadabbbadpwpw", "bbadabb"},
}
