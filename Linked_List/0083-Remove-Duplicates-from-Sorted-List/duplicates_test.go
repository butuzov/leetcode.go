package duplicates

import (
	"testing"
)

/*******************************************************************************
  Tests
*******************************************************************************/
func TestDeleteDuplicates(t *testing.T) {
	for _, test := range TestCases {
		actual := deleteDuplicates(test.input)
		for {
			if actual.Val != test.expected.Val {
				t.Errorf(MessageError, test.input, test.expected, actual)
				break
			}

			if actual.Next == test.expected.Next {
				t.Logf(MessageOk, test.input, test.expected)
				break
			}

			if (actual.Next == nil && test.expected.Next != nil) || (test.expected.Next == nil && actual.Next != nil) {
				t.Errorf(MessageError, test.input, test.expected, actual)
				break
			}

			actual = actual.Next
			test.expected = test.expected.Next
		}

	}
}

func BenchmarkDeleteDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			deleteDuplicates(test.input)
		}
	}
}
