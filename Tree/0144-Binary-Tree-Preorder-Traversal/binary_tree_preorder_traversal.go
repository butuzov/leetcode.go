package main

/*******************************************************************************
  Problem Solution
*******************************************************************************/

func preorderTraversalRecursive(root *TreeNode) []int {
	var results = []int{}
	if root == nil {
		return results
	}

	var preOrder func(*TreeNode)
	preOrder = func(t *TreeNode) {
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

func preorderTraversal(root *TreeNode) []int {
	var results = []int{}
	var stack = []*TreeNode{}
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
