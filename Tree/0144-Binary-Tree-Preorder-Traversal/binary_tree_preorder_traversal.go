package main

import "github.com/butuzov/leetcode.go/pkg/binarytree"

/*******************************************************************************
  Problem Solution
*******************************************************************************/

func preorderTraversalRecursive(root *binarytree.TreeNode) []int {
	var results = []int{}
	if root == nil {
		return results
	}

	var preOrder func(*binarytree.TreeNode)
	preOrder = func(t *binarytree.TreeNode) {
		if t == nil {
			return
		}

		results = append(results, t.Val)
		preOrder(t.Left)
		preOrder(t.Right)
	}

	preOrder(root)
	return results
}

func preorderTraversal(root *binarytree.TreeNode) []int {
	var results = []int{}
	var stack = []*binarytree.TreeNode{}
	var node = root

	for node != nil || len(stack) > 0 {

		if node != nil {
			results = append(results, node.Val)
		}

		if node.Left == nil && node.Right == nil && len(stack) == 0 {
			break
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if len(stack) > 0 {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

	}

	return results
}
