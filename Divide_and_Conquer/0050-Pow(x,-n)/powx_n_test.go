package main
import (
	"fmt"
	"testing"
)

var MsgError = "Fail: Input(x=(%f),n=(%d)): Expected(%f) vs Returend(%f)"
var MsgOk = "OK: Input(x=(%f),n=(%d)): as Expected(%f)"

func TestMyPow(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		test := test

		t.Run(fmt.Sprintf("MyPow(%d)", i), func(t *testing.T) {
			res := myPow(test.x, test.n)

			// reflect.DeepEqual(res, test.expected)
			if test.expected-res > 0.0001 {
				t.Errorf(MsgError, test.x, test.n, test.expected, res)
				return
			}
			t.Logf(MsgOk, test.x, test.n, test.expected)
		})
	}
}

func BenchmarkMyPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			myPow(test.x, test.n)
		}
	}
}
