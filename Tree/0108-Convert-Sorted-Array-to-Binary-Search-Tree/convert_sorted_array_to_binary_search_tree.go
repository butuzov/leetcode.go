package main

import (
	"fmt"

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

func sortedArrayToBST(nums []int) *binarytree.TreeNode {

	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2

	return &binarytree.TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}

}

func main() {

	// tree := binarytree.NewTree([]string{"3", "1l", "0l", "p", "2r", "p", "p", "5r", "4l"})
	// fmt.Println(tree)
	// fmt.Println(sortedArrayToBST([]int{0, 1, 2, 3, 4, 5}))

	tree := binarytree.NewTree([]string{"0", "-3l", "-10l", "p", "p", "9r", "5l"})
	fmt.Println(tree)
	fmt.Println(sortedArrayToBST([]int{-10, -3, 0, 5, 9}))

}
