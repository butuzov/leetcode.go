package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	croakOfFrogs string
	expected     int
}{
	{"croakcroak", 1},
	{"crcoakroak", 2},
	{"croakcrook", -1},
	{"croakcroa", -1},
}
