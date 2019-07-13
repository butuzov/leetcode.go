package main

/*******************************************************************************
  Problem Solution
*******************************************************************************/

func isValidBST(root *TreeNode) bool {
	var values = []int{}
	var traversal func(*TreeNode, *[]int)
	traversal = func(r *TreeNode, v *[]int) {
		if r == nil {
			return
		}

		traversal(r.Left, v)
		*v = append(*v, r.Val)
		traversal(r.Right, v)

	}

	traversal(root, &values)
	for i := 1; i < len(values); i++ {
		if values[i-1] >= values[i] {
			return false
		}
	}

	return true
}

// ------

func isValidBSTEveryNode(root *TreeNode) bool {

	var validBinTree func(*TreeNode, *[]int, bool) bool

	validBinTree = func(tree *TreeNode, s *[]int, left bool) bool {

		if tree == nil {
			return true
		}

		if tree.Left != nil && tree.Left.Val >= tree.Val {
			return false
		}

		if tree.Right != nil && tree.Right.Val <= tree.Val {
			return false
		}

		*s = append(*s, tree.Val)

		if !validBinTree(tree.Left, s, true) || !validBinTree(tree.Right, s, false) {
			return false
		}

		*s = (*s)[:len(*s)-1]

		if len(*s) < 2 {
			return true
		}
		for i := 0; i < len(*s)-1; i++ {

			// subparent is left
			if (*s)[i] > (*s)[i+1] {
				if tree.Val >= (*s)[i] {
					return false
				}
				continue
			}

			// subparent is right branch
			if (*s)[i] < (*s)[i+1] {
				if tree.Val <= (*s)[i] {
					return false
				}
			}
		}

		return true
	}

	var stack = []int{}
	return validBinTree(root, &stack, true)
}
