package main

import (
	"github.com/butuzov/leetcode.go/pkg/binarytree"
	"github.com/butuzov/leetcode.go/pkg/linkedlist"
)

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *linkedlist.ListNode) *binarytree.TreeNode {

	if head == nil {
		return nil
	}

	var length int
	var total int

	for node := head; node != nil; node = node.Next {
		total++
	}

	var (
		pre  *linkedlist.ListNode = head
		post *linkedlist.ListNode
		cur  *linkedlist.ListNode
		val  int
	)

	// spliting list into groups
	mid := total / 2

	for node := head; node != nil; node = node.Next {
		if mid == length {
			val = node.Val
			post = node.Next

			if pre == node {
				pre = nil
			}

			if cur != nil {
				cur.Next = nil
			}
			break
		}
		length++
		cur = node
	}

	return &binarytree.TreeNode{
		Val:   val,
		Left:  sortedListToBST(pre),
		Right: sortedListToBST(post),
	}
}
