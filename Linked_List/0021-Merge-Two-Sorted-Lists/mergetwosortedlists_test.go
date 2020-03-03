package main

import (
	"testing"

	"github.com/butuzov/leetcode.go/pkg/linkedlist"
)

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Inputs (%s) and (%s): Expected(%s) vs Returend(%s)"
var MessageOk = "OK: Inputs (%s) and (%s) as Expected(%s)"

var TestCases = []struct {
	input1   *linkedlist.ListNode
	input2   *linkedlist.ListNode
	expected *linkedlist.ListNode
}{
	{
		linkedlist.MakeList([]int{1}...),
		linkedlist.MakeList([]int{2}...),
		linkedlist.MakeList([]int{1, 2}...),
	},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestMergeTwoLists(t *testing.T) {
	for _, test := range TestCases {
		actual := mergeTwoLists(test.input1, test.input2)
		var preservedActual = actual
		var preservedExpected = test.expected
		for {
			if actual.Val != test.expected.Val {
				t.Errorf(MessageError, test.input1, test.input2, preservedExpected, preservedActual)
				break
			}

			if actual.Next == test.expected.Next {
				t.Logf(MessageOk, test.input1, test.input2, preservedExpected)
				break
			}

			if (actual.Next == nil && test.expected.Next != nil) || (test.expected.Next == nil && actual.Next != nil) {
				t.Errorf(MessageError, test.input1, test.input2, preservedExpected, preservedActual)
				break
			}

			actual = actual.Next
			test.expected = test.expected.Next
		}
	}
}

func BenchmarkMergeTwoLists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			mergeTwoLists(test.input1, test.input2)
		}
	}
}
