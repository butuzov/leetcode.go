package rotatelist

func rotateRight(list *ListNode, k int) *ListNode {
	if list == nil {
		return nil
	}

	var tmp = list
	var len = 1

	for tmp.Next != nil {
		tmp = tmp.Next
		len++
	}

	if k > len {
		k = k % len
	}

	if len == 1 || len == k || k == 0 {
		return list
	}

	var head = list
	var prev *ListNode
	var splitPoint = len - k

	// find the split point
	for i := 1; i <= splitPoint; i++ {
		prev = head
		head = head.Next
	}

	// seal tail
	prev.Next = nil

	// move part to begining
	tmp = head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = list

	return head
}
