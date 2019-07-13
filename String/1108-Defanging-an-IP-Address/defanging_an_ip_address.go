package main

import (
	"bytes"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func defangIPaddr(address string) string {
	var b bytes.Buffer
	b.Grow(len(address) + 6)
	for i := range address {
		if address[i] == 46 {
			b.Write([]byte{91, 46, 93})
			continue
		}
		b.WriteByte(address[i])
	}

	return b.String()
}
