package raindrops

import (
	"bytes"
	"fmt"
)

// Convert drops to "sound"?.
func Convert(drops int) string {

	buf := bytes.NewBuffer([]byte{})

	// This is pretty stupid way...
	// but its a simpiest way as well.
	if drops%3 == 0 {
		buf.WriteString("Pling")
	}
	if drops%5 == 0 {
		buf.WriteString("Plang")
	}
	if drops%7 == 0 {
		buf.WriteString("Plong")
	}

	if buf.Len() == 0 {
		buf.WriteString(fmt.Sprintf("%d", drops))
	}

	return buf.String()
}
