package main

import (
	"fmt"
	"strings"

	"github.com/butuzov/leetcode.go/pkg/binarytree"
)

func main() {
	// for i := range TestCases {
	// 	fmt.Println(TestCases[i].root, isValidBST(TestCases[i].root))
	// 	fmt.Println()
	// }
	// t := newTree(strings.Split("50|25l|12l|6l|p|14r|p|p|37r|15l|p|45r|p|p|p|75r|63l|58l|p|72r|p|p|87r|85l|p|98r|93l|p|99r", "|"))
	t := binarytree.NewTree(strings.Split("3|1l|0l|p|2r|3r|p|p|p|5r|4l|p|6r", "|"))

	fmt.Println(t)

	fmt.Println(isValidBST(t))
}
