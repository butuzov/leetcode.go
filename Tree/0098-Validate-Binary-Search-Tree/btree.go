package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// TreeNode represent of the binary tree nodes.
type TreeNode struct {
	Val    int
	parent *TreeNode
	Left   *TreeNode
	Right  *TreeNode
}

// ------- Creation ------------------------------------------------------------

// Node creation.
func Node(v int) *TreeNode {
	return &TreeNode{Val: v}
}

// newTree as set of instructions
// Example:
// 		newTree([]string{"1", "2l", "4l", "p", "5r", "p", "p", "3r"})
func newTree(instructions []string) *TreeNode {
	if len(instructions) == 0 {
		return nil
	}

	var node *TreeNode
	var root *TreeNode
	val, _ := strconv.Atoi(instructions[0])
	root = Node(val)
	node = root

	for _, instr := range instructions[1:] {
		val, _ := strconv.Atoi(string(instr[0 : len(instr)-1]))
		switch instr[len(instr)-1] {
		case 'l':
			node = node.InsertLeft(val)
		case 'r':
			node = node.InsertRight(val)
		case 'p':
			node = node.Parent()
		}
	}

	return root
}

// makeNewTree will create new Tree from the variatic call.
func makeNewTree(nums ...int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var tree = Node(nums[0])
	for _, i := range nums[1:] {
		tree.Insert(i)
	}
	return tree
}

// Insert inserts a node into binary tre
func (t *TreeNode) Insert(value int) *TreeNode {
	if t.Val < value {
		if t.Right != nil {
			return t.Right.Insert(value)
		}

		t.Right = &TreeNode{Val: value}
		return t.Right

	}

	if t.Left != nil {
		return t.Left.Insert(value)
	}
	t.Left = &TreeNode{Val: value}
	return t.Left
}

// Parent will return parent node. It's not suppose to be used in any
// leetcode test, as method that not suppose to be there.
func (t *TreeNode) Parent() *TreeNode {
	return t.parent
}

// InsertLeft will insert value to left (no matter is there is any data)
func (t *TreeNode) InsertLeft(i int) *TreeNode {
	t.Left = &TreeNode{Val: i, parent: t}
	return t.Left
}

// InsertRight will insert value to left (no matter is there is any data)
func (t *TreeNode) InsertRight(i int) *TreeNode {
	t.Right = &TreeNode{Val: i, parent: t}
	return t.Right
}

// ------- Traversing ----------------------------------------------------------

// InOrder Traversal
func (t *TreeNode) InOrder() string {
	if t == nil {
		return ""
	}
	var s = t.traversalInOrder()
	return s[2:]
}

// traversalInOrder is helper method that travers bt in order (left/root/right).
func (t *TreeNode) traversalInOrder() string {
	var b bytes.Buffer

	if t == nil {
		return ""
	}

	b.WriteString(t.Left.traversalInOrder())
	fmt.Fprintf(&b, "->%d", t.Val)
	b.WriteString(t.Right.traversalInOrder())

	return b.String()
}

// PostOrder traversal printer
func (t *TreeNode) PostOrder() string {

	if t == nil {
		return ""
	}
	var s = t.traversalPostOrder()
	return s[2:]
}

// traversalPostOrder is helper that travers bt in order (left/right/root)
func (t *TreeNode) traversalPostOrder() string {
	if t == nil {
		return ""
	}

	var b bytes.Buffer

	b.WriteString(t.Left.traversalPostOrder())
	b.WriteString(t.Right.traversalPostOrder())
	fmt.Fprintf(&b, "->%d", t.Val)

	return b.String()
}

// PreOrder traversal printer
func (t *TreeNode) PreOrder() string {
	if t == nil {
		return ""
	}
	var s = t.traversalPreOrder()
	return s[2:]
}

// traversalPreOrder is helper that travers bt in order (root/left/right)
func (t *TreeNode) traversalPreOrder() string {
	if t == nil {
		return ""
	}

	var b bytes.Buffer

	fmt.Fprintf(&b, "->%d", t.Val)
	b.WriteString(t.Left.traversalPreOrder())
	b.WriteString(t.Right.traversalPreOrder())

	return b.String()
}

// ------- Debug ---------------------------------------------------------------

// String helps to debug binary tree representation.
func (t *TreeNode) String() string {

	var BT_HEIGHT = t.Height()
	var BT_WIDTH = 80 // arbitrary number of cols per matrix output.

	var tree func(*TreeNode, []byte, bool, int, int) int
	tree = func(t *TreeNode, canvas []byte, is_left bool, offset int, depth int) int {
		if t == nil {
			return 0
		}
		val := fmt.Sprintf("%d", t.Val)

		var left = tree(t.Left, canvas, true, offset, depth+1)
		var right = tree(t.Right, canvas, false, offset+left+len(val), depth+1)

		for i := 0; i < len(val); i++ {
			canvas[(BT_WIDTH*depth)+(offset+left+i)] = val[i]
		}

		var row = (depth - 1) * BT_WIDTH
		if depth > 0 && is_left {

			for i := 0; i < len(val)+right; i++ {
				canvas[row+(offset+left+len(val)/2+i)] = '-'
			}
			canvas[row+(offset+left+len(val)/2)] = '.'
		}

		if depth > 0 && !is_left {
			for i := 0; i < left+len(val); i++ {
				canvas[row+(offset-len(val)/2+i)] = '-'
			}
			canvas[row+(offset+left+len(val)/2)] = '.'
		}

		return left + len(val) + right
	}

	if t == nil {
		return ""
	}

	var b = make([]byte, BT_HEIGHT*BT_WIDTH)
	for i := 0; i < len(b); i++ {
		b[i] = ' '
	}
	tree(t, b, false, 0, 0)

	var canvas bytes.Buffer
	for i := 0; i < len(b); i += BT_WIDTH {
		canvas.Write(b[i : i+BT_WIDTH])
		canvas.WriteByte(10)
	}

	return canvas.String()

}

// Height give us a maximum depth of the Tree
func (t *TreeNode) Height() int {

	if t == nil {
		return 0
	}

	var (
		lHeight = t.Left.Height() + 1
		rHeight = t.Right.Height() + 1
	)

	if lHeight > rHeight {
		return lHeight
	}

	return rHeight
}
