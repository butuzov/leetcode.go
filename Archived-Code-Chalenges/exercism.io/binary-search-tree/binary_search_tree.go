package binarysearchtree

import (
	"fmt"
	"strconv"
)

// Basic SearchTreeData struct
type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// test function for easier testing
// cmd+r in vsc
func main() {
	bst := Bst(50)
	bst.Insert(25)
	bst.Insert(12)
	bst.Insert(37)
	bst.Insert(75)
	bst.Insert(53)
	bst.Insert(87)

	fmt.Println(bst.MapString(strconv.Itoa))

}

// Bst return BinaryTreeData struct with initialted node value.
func Bst(number int) *SearchTreeData {
	bst := new(SearchTreeData)
	bst.data = number
	return bst
}

// Insert method perform BinaryTree insert
// values less than node value (data) goes to left ( branch)
// bigger or equal goes to right branch.
// if branch isn't nil, we create new Bst and insert into branch.
func (bst *SearchTreeData) Insert(number int) {
	if bst.data < number {
		if bst.right != nil {
			bst.right.Insert(number)
		} else {
			bst.right = Bst(number)
		}
	}

	if bst.data >= number {
		if bst.left != nil {
			bst.left.Insert(number)
		} else {
			bst.left = Bst(number)
		}
	}
}

// Next two functions almost same and difference only in input callback and
// retunred results.., oh.. why to Go don't have generics.

// MapString - perform tree traversal and mapping values to string, with provided
// callback function.
func (bst *SearchTreeData) MapString(f func(int) string) []string {
	var out []string

	if bst.left != nil {
		for _, v := range bst.left.MapString(f) {
			out = append(out, v)
		}
	}

	out = append(out, f(bst.data))

	if bst.right != nil {
		for _, v := range bst.right.MapString(f) {
			out = append(out, v)
		}
	}

	return out
}

// MapInt - perform tree traversal and mapping values to int, with provided
// callback function.
func (bst *SearchTreeData) MapInt(f func(int) int) []int {
	var out []int

	if bst.left != nil {
		for _, v := range bst.left.MapInt(f) {
			out = append(out, v)
		}
	}

	out = append(out, f(bst.data))

	if bst.right != nil {
		for _, v := range bst.right.MapInt(f) {
			out = append(out, v)
		}
	}

	return out
}
