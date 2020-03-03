package main

import "github.com/butuzov/leetcode.go/pkg/linkedlist"

/*****************************************************************************
 * Solution
 ****************************************************************************/

// Solution.
func deleteDuplicates(head *linkedlist.ListNode) *linkedlist.ListNode {

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
