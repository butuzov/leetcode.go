package main
import (
	"testing"
)

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Inputs (%s) and (%s): Expected(%s) vs Returend(%s)"
var MessageOk = "OK: Inputs (%s) and (%s) as Expected(%s)"

var TestCases = []struct {
	input1   *ListNode
	input2   *ListNode
	expected *ListNode
}{
	{
		makeList([]int{1}...),
		makeList([]int{2}...),
		makeList([]int{1, 2}...),
	},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestMergeTwoLists(t *testing.T) {
	for _, test := range TestCases {
		actual := mergeTwoLists(test.input1, test.input2)
		var preservedActal = actual
		var preservedExpected = test.expected
		for {
			if actual.Val != test.expected.Val {
				t.Errorf(MessageError, test.input1, test.input2, preservedExpected, preservedActal)
				break
			}

			if actual.Next == test.expected.Next {
				t.Logf(MessageOk, test.input1, test.input2, preservedExpected)
				break
			}

			if (actual.Next == nil && test.expected.Next != nil) || (test.expected.Next == nil && actual.Next != nil) {
				t.Errorf(MessageError, test.input1, test.input2, preservedExpected, preservedActal)
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
