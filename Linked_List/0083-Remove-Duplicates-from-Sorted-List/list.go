package duplicates

import (
	"fmt"
	"strconv"
)

// ListNode is Linked List Node
type ListNode struct {
	Val  int
	Next *ListNode
}

// For a testing and development reasons we need to simplify
// Linked list creation.
func makeList(nums ...int) *ListNode {
	if len(nums) > 1 {
		return &ListNode{
			Val:  nums[0],
			Next: makeList(nums[1:]...),
		}
	}
	return &ListNode{Val: nums[0]}
}

// Stringer interface implementation for a nice linkedlist
// representation as string.
func (ln *ListNode) String() string {
	if ln.Next == nil {
		return strconv.Itoa(ln.Val)
	}
	return strconv.Itoa(ln.Val) + "->" + ln.Next.String()
}

func ExampleMakeLInkedList() {
	fmt.Println(makeList([]int{1, 2, 2, 3, 3, 3, 1, 1, 12, 3}...))
	// Output: 1->2->2->3->3->3->1->1->12->3
}
