package reachingpoints

import (
	"fmt"
	"testing"
)

var MsgError = "Fail: Input(sx=(%d),sy=(%d),tx=(%d),ty=(%d)): Expected(%d) vs Returend(%d)"
var MsgOk = "OK: Input(sx=(%d),sy=(%d),tx=(%d),ty=(%d)): as Expected(%d)"

func TestReachingPoints(t *testing.T) {
	t.Parallel()
	for _, test := range TestCases {
		test := test

		name := fmt.Sprintf("sx=(%d)sy=(%d) to tx=(%d),ty=(%d)", test.sx, test.sy, test.tx, test.ty)
		t.Run(name, func(t *testing.T) {
			res := reachingPoints(test.sx, test.sy, test.tx, test.ty)

			if res != test.expected {
				t.Errorf(MsgError, test.sx, test.sy, test.tx, test.ty, test.expected, res)
				return
			}
			t.Logf(MsgOk, test.sx, test.sy, test.tx, test.ty, test.expected)
		})
	}
}

func BenchmarkReachingPoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			reachingPoints(test.sx, test.sy, test.tx, test.ty)
		}
	}
}
