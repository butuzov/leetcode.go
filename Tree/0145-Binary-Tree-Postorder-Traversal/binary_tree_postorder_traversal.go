package main

import "github.com/butuzov/leetcode.go/pkg/binarytree"

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func postorderTraversalRecursive(root *binarytree.TreeNode) []int {
	var results = []int{}
	if root == nil {
		return results
	}

	var n = 0

	var postOrder func(*binarytree.TreeNode)
	postOrder = func(t *binarytree.TreeNode) {
		if t == nil {
			return
		}

		postOrder(t.Left)
		postOrder(t.Right)
		results = append(results, t.Val)
		n++
	}

	postOrder(root)
	return results
}

func postorderTraversal(root *binarytree.TreeNode) []int {
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

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if len(stack) > 0 {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}

	for i, j := 0, len(results)-1; i < j; i, j = i+1, j-1 {
		results[i], results[j] = results[j], results[i]
	}

	return results
}
