package main

import (
	"github.com/butuzov/leetcode.go/pkg/binarytree"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/

type BSTIteratorI struct {
	stack []*binarytree.TreeNode
	start int
}

func Constructor(root *binarytree.TreeNode) BSTIterator {
	var stack []*binarytree.TreeNode

	var populate func(*binarytree.TreeNode)
	populate = func(node *binarytree.TreeNode) {
		if node == nil {
			return
		}
		populate(node.Left)
		stack = append(stack, node)
		populate(node.Right)
	}
	populate(root)

	return BSTIterator{
		stack: stack,
		start: 0,
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	this.start++
	return this.stack[this.start-1].Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.start < len(this.stack)
}
