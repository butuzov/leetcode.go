package main

import (
	"testing"

	"github.com/butuzov/leetcode.go/pkg/binarytree"
	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	var iter = Constructor(binarytree.NewTree([]string{"7", "3l", "p", "15r", "9l", "p", "20r"}))
	assert.True(t, iter.HasNext())
	assert.Equal(t, 3, iter.Next())
	assert.Equal(t, 7, iter.Next())
	assert.True(t, iter.HasNext())
	assert.Equal(t, 9, iter.Next())
	assert.True(t, iter.HasNext())
	assert.Equal(t, 15, iter.Next())
	assert.True(t, iter.HasNext())
	assert.Equal(t, 20, iter.Next())
	assert.False(t, iter.HasNext())
}
