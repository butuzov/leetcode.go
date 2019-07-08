package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	equations []string
	expected  bool
}{
	{[]string{"a==b", "b!=a"}, false},
	{[]string{"b==a", "a==b"}, true},
	{[]string{"a==b", "b==c", "a==c"}, true},
	{[]string{"a==b", "b!=c", "c==a"}, false},
	{[]string{"c==c", "b==d", "x!=z"}, true},
	{[]string{"c!=c"}, false},
	{[]string{"b==b", "b==e", "e==c", "d!=e"}, true},
	{[]string{"b==b", "b==e", "e==c", "e!=d"}, true},
	{[]string{"b==b", "b==e", "c==e", "e!=d"}, true},
	{[]string{"c==c", "f!=a", "f==b", "b==c"}, true},
}
