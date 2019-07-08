package main

/*******************************************************************************
  Problem Solution
*******************************************************************************/

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Right == nil && root.Left == nil {
		return
	}

	// recursive call to fix left and right branches.
	flatten(root.Left)
	flatten(root.Right)

	// nothing to move right
	if root.Left == nil {
		return
	}

	// right branch become rightmost branch of the left
	right, left := root.Right, root.Left
	rightMost := left

	for rightMost != nil && rightMost.Right != nil {
		rightMost = rightMost.Right
	}

	rightMost.Right = right

	// left moving to old right
	root.Right = left

	// stoper
	root.Left = nil
}
