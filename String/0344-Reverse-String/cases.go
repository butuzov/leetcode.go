package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	text, expected []byte
}{
	{[]byte("hello"), []byte("olleh")},
	{[]byte("A man, a plan, a canal: Panama"), []byte("amanaP :lanac a ,nalp a ,nam A")},
}
