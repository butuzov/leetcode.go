package poweroffour

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	num      int
	expected bool
}{
	{4, true},
	{5, false},
	{8, false},
	{-16, false},
	{-1, false},
	{1, true},
}
