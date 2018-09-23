package duplicates

import (
	"strconv"
)

/*****************************************************************************
 * Helpers
 ****************************************************************************/

// CreateList will create linked list from numbers
func makeList(nums ...int) *ListNode {

	if len(nums) > 1 {
		return &ListNode{
			Val:  nums[0],
			Next: makeList(nums[1:]...),
		}
	}
	return &ListNode{Val: nums[0]}
}

// String is stringer interface implementation
func (ln *ListNode) String() string {
	if ln.Next == nil {
		return strconv.Itoa(ln.Val)
	}
	return strconv.Itoa(ln.Val) + "->" + ln.Next.String()
}

// Case definings truct of our tests.
type Case struct {
	list     *ListNode
	expected *ListNode
}

// makeTestCase - makes a Test Case =)
func makeTestCase(list, expected *ListNode) Case {
	return Case{
		list:     list,
		expected: expected,
	}
}

/*****************************************************************************
 * Solution
 ****************************************************************************/

// ListNode is Linked List Node
type ListNode struct {
	Val  int
	Next *ListNode
}

// Solution.
func deleteDuplicates(head *ListNode) *ListNode {

	if head != nil {
		current := head

		for current.Next != nil {
			if current.Val == current.Next.Val {
				current.Next = current.Next.Next
			} else {
				current = current.Next
			}
		}
	}

	return head
}
