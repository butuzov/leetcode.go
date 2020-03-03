package main

import (
	"github.com/butuzov/leetcode.go/pkg/binarytree"
)

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
func isSameTree(p *binarytree.TreeNode, q *binarytree.TreeNode) bool {

	// boath are nil
	if p == q {
		return true
	} else if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
