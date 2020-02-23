package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ["LRUCache","put","put","get","put","put","get"]
// [[2],[2,1],[2,2],[2],[1,1],[4,1],[2]]

func TestOne(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	assert.Equal(t, cache.Get(1), 1)
	cache.Put(3, 3)
	assert.Equal(t, cache.Get(2), -1)
	cache.Put(4, 4)
	assert.Equal(t, cache.Get(1), -1)
	assert.Equal(t, cache.Get(3), 3)
	assert.Equal(t, cache.Get(4), 4)
}

func TestTwo(t *testing.T) {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	assert.Equal(t, cache.Get(2), 2)
	cache.Put(1, 1)
	cache.Put(4, 1)
	assert.Equal(t, cache.Get(2), -1)
}

func TestThree(t *testing.T) {
	cache := Constructor(3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)
	assert.Equal(t, cache.Get(4), 4)
	assert.Equal(t, cache.Get(3), 3)
	assert.Equal(t, cache.Get(2), 2)
	assert.Equal(t, cache.Get(1), -1)
	cache.Put(5, 5)
	assert.Equal(t, cache.Get(1), -1)
	assert.Equal(t, cache.Get(2), 2)
	assert.Equal(t, cache.Get(3), 3)
	assert.Equal(t, cache.Get(4), -1)
	assert.Equal(t, cache.Get(5), 5)
}
