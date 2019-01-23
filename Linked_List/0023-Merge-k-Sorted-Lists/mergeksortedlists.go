package mergeksortedlists

import (
	"fmt"
	"sort"
	"strconv"
)

// ListNode is Linked List Node
type ListNode struct {
	Val  int
	Next *ListNode
}

// Sorts K/N linked lists into one.
func mergeKLists(linkedLists []*ListNode) *ListNode {

	var node *ListNode
	var head *ListNode
	var lists []*ListNode

	for _, l := range linkedLists {
		if l != nil {
			lists = append(lists, l)
		}
	}

	switch len(lists) {
	case 0:
		return head
	case 1:
		return lists[0]
	}

	sort.Slice(lists, func(i, j int) bool {
		return lists[i].Val < lists[j].Val
	})

	for {

		if len(lists) > 1 && lists[0].Val > lists[1].Val {
			sort.Slice(lists, func(i, j int) bool {
				return lists[i].Val < lists[j].Val
			})
		}

		link := &ListNode{Val: lists[0].Val}

		if node == nil {
			head = link
			node = link
		} else {
			node.Next = link
			node = link
		}

		if lists[0].Next != nil {
			lists[0] = lists[0].Next
		} else if len(lists) > 1 {
			lists = append(lists[:0], lists[1:]...)
		} else {
			break
		}
	}

	return head
}

// Lists type represents a
type Lists []*ListNode

func (l Lists) Len() int           { return len(l) }
func (l Lists) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l Lists) Less(i, j int) bool { return l[i].Val < l[j].Val }

// For a testing and development reasons we need to siplify
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
