package main

import (
	"testing"
)

/*******************************************************************************
  Tests
*******************************************************************************/
func TestAddTwoNumbers(t *testing.T) {
	for _, test := range TestCases {
		actual := addTwoNumbers(test.llOne, test.llTwo)
		for {
			if actual.Val != test.llResult.Val {
				t.Errorf(MessageError, test.llOne, test.llTwo, test.llResult, actual)
				break
			}

			if actual.Next == test.llResult.Next {
				t.Logf(MessageOk, test.llOne, test.llTwo, test.llResult)
				break
			}

			if (actual.Next == nil && test.llResult.Next != nil) || (test.llResult.Next == nil && actual.Next != nil) {
				t.Errorf(MessageError, test.llOne, test.llTwo, test.llResult, actual)
				break
			}

			actual = actual.Next
			test.llResult = test.llResult.Next
		}

	}
}

func BenchmarkTestAddTwoNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			addTwoNumbers(test.llOne, test.llTwo)
		}
	}
}
