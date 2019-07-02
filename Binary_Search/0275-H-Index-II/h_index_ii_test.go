package hindexii

import (
	"fmt"
	"testing"
)

var MsgError = "Fail: Input(citations=(%+v)): Expected(%+v) vs Returend(%+v)"
var MsgOk = "OK: Input(citations=(%+v)): as Expected(%+v)"

func TestHIndex(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		test := test

		t.Run(fmt.Sprintf("H-Index(%d)", i), func(t *testing.T) {
			res := hIndex(test.citations)

			if res != test.expected {
				t.Errorf(MsgError, test.citations, test.expected, res)
				return
			}
			t.Logf(MsgOk, test.citations, test.expected)
		})
	}
}

func BenchmarkHIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			hIndex(test.citations)
		}
	}
}
