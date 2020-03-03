package main

import (
	"fmt"

	"github.com/butuzov/leetcode.go/pkg/binarytree"
)

func main() {

	var tree = binarytree.NewTree([]string{"1", "2l", "3l", "p", "4r", "p", "p", "5r", "6r"})
	// var tree = newTree([]string{"2", "3l", "p", "4r"})
	fmt.Println(tree)
	fmt.Println("Traversal In_Order", tree.InOrder())
	fmt.Println("Traversal Post_Order", tree.PostOrder())
	fmt.Println("Traversal Pre_Order", tree.PreOrder())

	flatten(tree)

	fmt.Println(tree)
	fmt.Println("Traversal In_Order", tree.InOrder())
	fmt.Println("Traversal Post_Order", tree.PostOrder())
	fmt.Println("Traversal Pre_Order", tree.PreOrder())

}
