package main
import (
	"fmt"
	"testing"
)

var MsgError = "Fail: Input(num=(%+v)): Expected(%+v) vs Returend(%+v)"
var MsgOk = "OK: Input(num=(%+v)): as Expected(%+v)"

func TestIsPowerOfFour(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		test := test
		t.Run(fmt.Sprintf("Power_4(%d)", i), func(t *testing.T) {
			res := isPowerOfFour(test.num)

			if res != test.expected {
				t.Errorf(MsgError, test.num, test.expected, res)
				return
			}
			t.Logf(MsgOk, test.num, test.expected)
		})
	}
}

func BenchmarkIsPowerOfFour(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			isPowerOfFour(test.num)
		}
	}
}
