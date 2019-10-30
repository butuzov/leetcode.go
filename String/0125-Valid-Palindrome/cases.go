package main

/*******************************************************************************
  TestTable
*******************************************************************************/

import (
	"io/ioutil"
	"log"
)

type Case struct {
	s        string
	expected bool
}

var TestCases = []Case{
	{"A man, a plan, a canal: Panama", true},
	{"race a car", false},
	{":;", true},
}

func init() {
	b, e := ioutil.ReadFile("really_long_palindrom.txt")
	if e != nil {
		log.Fatal(e)
	}

	TestCases = append(TestCases, Case{string(b), true})
}
