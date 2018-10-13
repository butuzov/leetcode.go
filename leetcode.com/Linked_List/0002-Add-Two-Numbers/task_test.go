package addtwonumbers

import (
	"strconv"
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	if ln.Next == nil {
		return strconv.Itoa(ln.Val)
	}
	return strconv.Itoa(ln.Val) + "->" + ln.Next.String()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var numbers []int
	var extra int
	var number int

	for {
		number = 0
		var done bool

		if l1 != nil {
			number += l1.Val
			l1 = l1.Next
			done = true
		}

		if l2 != nil {
			number += l2.Val
			l2 = l2.Next
			done = true
		}

		if extra > 0 {
			number += extra
			done = true
		}

		if done == false {
			break
		}

		numbers = append(numbers, number%10)

		if number >= 10 {
			extra = 1
		} else {
			extra = 0
		}
	}

	return makeList(numbers...)
}

func makeList(nums ...int) *ListNode {

	if len(nums) > 1 {
		return &ListNode{
			Val:  nums[0],
			Next: makeList(nums[1:]...),
		}
	}
	return &ListNode{Val: nums[0]}
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
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

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			addTwoNumbers(test.llOne, test.llTwo)
		}
	}
}
