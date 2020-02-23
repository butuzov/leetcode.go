package main

import (
	"bytes"
	"container/list"
	"fmt"
)

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

type LRUCache struct {
	size int
	list *list.List
	mapa map[int]*list.Element
}

type Item struct {
	Key   int
	Value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		size: capacity,
		list: list.New(),
		mapa: make(map[int]*list.Element),
	}
}

// retrive value for key, -1 if key not found.
func (this *LRUCache) Get(key int) int {
	var ok bool
	var el *list.Element

	if el, ok = this.mapa[key]; !ok {
		return -1
	}

	// Get Value
	this.list.MoveToFront(el)

	return el.Value.(Item).Value
}

// Put: adds value to LRU cache or updates if value in cache.
func (this *LRUCache) Put(key int, value int) {

	if el, ok := this.mapa[key]; ok {
		var item = el.Value.(Item)
		item.Value = value
		el.Value = item
		this.list.MoveToFront(el)
		return
	}

	this.mapa[key] = this.list.PushFront(Item{Key: key, Value: value})

	if this.list.Len() > this.size {
		var el = this.list.Back()
		var item = el.Value.(Item)
		delete(this.mapa, item.Key)
		this.list.Remove(el)
	}
}

func (this LRUCache) String() string {

	if this.list.Len() == 0 {
		return "[]"
	}

	var buf bytes.Buffer
	for e := this.list.Front(); e != nil; e = e.Next() {
		fmt.Fprintf(&buf, "%d<->", e.Value.(Item).Value)
	}
	var str = buf.String()
	return fmt.Sprintf("[%s]", str[:len(str)-3])
}
