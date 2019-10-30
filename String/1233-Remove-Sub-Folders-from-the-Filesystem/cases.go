package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	folder   []string
	expected []string
}{
	{[]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}, []string{"/a", "/c/d", "/c/f"}},
	{[]string{"/a", "/a/b/c", "/a/b/d"}, []string{"/a"}},
	{[]string{"/a/b/c", "/a/b/ca", "/a/b/d"}, []string{"/a/b/c", "/a/b/ca", "/a/b/d"}},
	{[]string{"/ah/al/am", "/ah/al"}, []string{"/ah/al"}},
}
