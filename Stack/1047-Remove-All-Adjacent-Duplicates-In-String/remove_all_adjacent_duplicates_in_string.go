package main

import (
	"bytes"
	"container/list"
)

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func removeDuplicatesList(input string) string {

	var stack = list.New()

	for i, _ := range input {
		if stack.Len() > 0 && stack.Back().Value.(byte) == input[i] {
			stack.Remove(stack.Back())
			continue
		}
		stack.PushBack(input[i])
	}

	// final part output
	var output bytes.Buffer
	output.Grow(len(input))

	for e := stack.Front(); e != nil; e = e.Next() {
		output.WriteByte(e.Value.(byte))
	}
	return output.String()
}

func removeDuplicates(input string) string {

	var stack = make([]byte, 0, len(input))

	for i, _ := range input {
		size := len(stack)
		if size > 0 && stack[size-1] == input[i] {
			stack = stack[:size-1]
			continue
		}
		stack = append(stack, input[i])
	}

	return string(stack)
}
