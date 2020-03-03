package main

import (
	"sort"

	"github.com/butuzov/leetcode.go/pkg/linkedlist"
)

// Sorts K/N linked lists into one.
func mergeKLists(linkedLists []*linkedlist.ListNode) *linkedlist.ListNode {

	var node *linkedlist.ListNode
	var head *linkedlist.ListNode
	var lists []*linkedlist.ListNode

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

		link := &linkedlist.ListNode{Val: lists[0].Val}

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
type Lists []*linkedlist.ListNode

func (l Lists) Len() int           { return len(l) }
func (l Lists) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l Lists) Less(i, j int) bool { return l[i].Val < l[j].Val }
