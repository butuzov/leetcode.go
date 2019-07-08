package main

/*******************************************************************************
  Problem Solution
*******************************************************************************/

func inorderTraversalRecursive(root *TreeNode) []int {
	var results = []int{}
	if root == nil {
		return results
	}

	var inOrder func(*TreeNode)
	inOrder = func(t *TreeNode) {
		if t == nil {
			return
		}

		inOrder(t.Left)
		results = append(results, t.Val)
		inOrder(t.Right)
	}

	inOrder(root)
	return results
}

func inorderTraversal(root *TreeNode) []int {
	var results = []int{}
	var stack = []*TreeNode{}
	var node = root

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		if len(stack) > 0 {
			node = stack[len(stack)-1]
			results = append(results, node.Val)
			stack = stack[:len(stack)-1]
			node = node.Right
		}

	}

	return results
}
