package transpose

import (
	"bytes"
)

func Transpose(in []string) []string {
	var bufs = make([]bytes.Buffer, 0)

	for col, str := range in {
		for row := 0; row < len(str); row++ {
			if len(bufs)-1 < row {
				bufs = append(bufs, *bytes.NewBuffer([]byte{}))
			}

			for bufs[row].Len() < col {
				bufs[row].WriteByte(32)
			}

			bufs[row].WriteByte(str[row])
		}
	}

	var out = make([]string, 0)
	for _, buf := range bufs {
		out = append(out, buf.String())
	}
	return out
}
