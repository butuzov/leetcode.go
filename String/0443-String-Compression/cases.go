package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	chars    []byte
	expected []byte
	len      int
}{
	{[]byte("a"), []byte("a"), 1},
	{[]byte("aabbccc"), []byte("a2b2c3"), 6},
	{[]byte("abbbbbbbbbbbb"), []byte("ab12"), 4},
	{[]byte("abb"), []byte("ab2"), 3},
	{[]byte("abbb"), []byte("ab3"), 3},
	{[]byte("abbbccc"), []byte("ab3c3"), 5},
	{[]byte("aabbbccc"), []byte("a2b3c3"), 6},
	{[]byte("aabccc"), []byte("a2bc3"), 5},
	{[]byte("abc"), []byte("abc"), 3},
	{[]byte("aabc"), []byte("a2bc"), 4},
	{[]byte("abbc"), []byte("ab2c"), 4},
	{[]byte("aabbbbbbbbbbbbb"), []byte("a2b13"), 5},
	{[]byte("abcdefggggggggggggabc"), []byte("abcdefg12abc"), 12},
}
