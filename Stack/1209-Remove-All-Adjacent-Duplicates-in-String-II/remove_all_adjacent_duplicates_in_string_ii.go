package main

import (
	"bytes"
	"container/list"
	"fmt"
)

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

type KeyVal struct {
	val   byte
	count int
}

func (kv *KeyVal) Increase() {
	kv.count += 1
}

func (kv KeyVal) String() string {
	return fmt.Sprintf("%c[%d]", kv.val, kv.count)
}

func (kv1 KeyVal) SameAs(kv2 KeyVal) bool {
	return kv1.val == kv2.val
}

func (kv KeyVal) Count() int {
	return kv.count
}

func (kv KeyVal) Val() byte {
	return kv.val
}

//  removeDuplicates will use container/list to keep track of the
// previous entry
func removeDuplicates(input string, k int) string {

	l := list.New()
	for i := range input {
		current := KeyVal{val: input[i], count: 1}

		if l.Len() == 0 {
			l.PushBack(current)
			continue
		}

		last := l.Back().Value.(KeyVal)
		if last.SameAs(current) {
			last.Increase()

			if last.Count() != k {
				l.Back().Value = last
				continue
			}

			_ = l.Remove(l.Back())
			continue
		}

		l.PushBack(current)
	}

	// final part output
	var output bytes.Buffer
	output.Grow(len(input))

	for e := l.Front(); e != nil; e = e.Next() {
		el := e.Value.(KeyVal)
		for i := 0; i < el.Count(); i++ {
			output.WriteByte(el.val)
		}
	}
	return output.String()
}
