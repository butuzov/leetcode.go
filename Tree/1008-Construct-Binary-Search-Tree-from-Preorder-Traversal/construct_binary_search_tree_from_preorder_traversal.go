package main

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	var tree = &TreeNode{Val: preorder[0]}

	if len(preorder) == 1 {
		return tree
	}

	var insert func(*TreeNode, int)

	insert = func(t *TreeNode, v int) {

		// left part
		if t.Val > v {
			if t.Left != nil {
				insert(t.Left, v)
				return
			}

			t.Left = &TreeNode{Val: v}
			return
		}

		// right part

		if t.Right != nil {
			insert(t.Right, v)
			return
		}

		t.Right = &TreeNode{Val: v}
		return
	}

	for _, i := range preorder[1:] {
		insert(tree, i)
	}

	return tree
}
