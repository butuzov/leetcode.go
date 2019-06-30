package duplicates

/*****************************************************************************
 * Solution
 ****************************************************************************/

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
