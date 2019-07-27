package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	S        string
	expected string
}{
	{"ab-cd", "dc-ba"},
	{"a-bC-dEf-ghIj", "j-Ih-gfE-dCba"},
	{"Test1ng-Leet=code-Q!", "Qedo1ct-eeLg=ntse-T!"},
}
