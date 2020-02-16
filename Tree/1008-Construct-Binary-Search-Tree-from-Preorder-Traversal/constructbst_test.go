package main

import (
	"testing"

	"github.com/butuzov/leetcode.go/pkg/binarytree"
	"github.com/stretchr/testify/assert"
)

func TestPreOrder(t *testing.T) {

	assert.Equal(t, binarytree.MakeNewTree([]int{8, 5, 1, 7, 10, 12}...).InOrder(), bstFromPreorder([]int{8, 5, 1, 7, 10, 12}).InOrder())
}
