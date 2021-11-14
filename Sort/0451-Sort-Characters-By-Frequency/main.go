package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// dude, write some code...
	fmt.Printf(frequencySort("tree"))
}

func frequencySort(s string) string {
	type kv struct {
		char  byte
		count int
	}

	var counter []kv
	charIndex := map[byte]int{}
	for i := range s {
		if idx, ok := charIndex[s[i]]; ok {
			c := counter[idx]
			c.count++
			counter[idx] = c
			continue
		}

		counter = append(counter, kv{char: s[i], count: 1})
		charIndex[s[i]] = len(counter) - 1
	}

	sort.Slice(counter, func(i, j int) bool {
		return counter[i].count > counter[j].count
	})

	var sb strings.Builder
	for _, kv := range counter {
		sb.WriteString(strings.Repeat(string(kv.char), kv.count))
	}

	return sb.String()
}
