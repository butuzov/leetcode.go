package main

import (
	"github.com/butuzov/leetcode.go/pkg/binarytree"
)

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func bstFromPreorder(preorder []int) *binarytree.TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	var tree = &binarytree.TreeNode{Val: preorder[0]}

	if len(preorder) == 1 {
		return tree
	}

	var insert func(*binarytree.TreeNode, int)

	insert = func(t *binarytree.TreeNode, v int) {

		// left part
		if t.Val > v {
			if t.Left != nil {
				insert(t.Left, v)
				return
			}

			t.Left = &binarytree.TreeNode{Val: v}
			return
		}

		// right part

		if t.Right != nil {
			insert(t.Right, v)
			return
		}

		t.Right = &binarytree.TreeNode{Val: v}
		return
	}

	for _, i := range preorder[1:] {
		insert(tree, i)
	}

	return tree
}
