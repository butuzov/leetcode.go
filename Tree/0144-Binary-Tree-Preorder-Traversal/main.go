package main

import (
	"fmt"
	"strings"

	"github.com/butuzov/leetcode.go/pkg/binarytree"
)

func main() {
	t := binarytree.NewTree(strings.Split("25|12l|10l|4l|1l|p|p|p|5r|p|p|4r|0r|31r|100r", "|"))

	fmt.Println(t)
	fmt.Println("PreOrder_iStringer", t.PreOrder())
	fmt.Println("PreOrder_Recursive", preorderTraversalRecursive(t))
	fmt.Println("PreOrder_Naive    ", preorderTraversal(t))
}
