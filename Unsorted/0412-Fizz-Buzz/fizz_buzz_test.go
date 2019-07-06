package main
import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var MsgError = "Fail: Input(n=(%+v)): Expected(%+v) vs Returend(%+v)"
var MsgOk = "OK: Input(n=(%+v)): as Expected(%+v)"

func TestFizzBuzz(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		test := test

		t.Run(fmt.Sprintf("FizzBuzz(%d)", i), func(t *testing.T) {
			assert.Equalf(t, test.expected, fizzBuzz(test.n), "Input(n=(%+v))", test.n)
		})
	}
}

func BenchmarkFizzBuzz(b *testing.B) {

	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				fizzBuzz(test.n)
			}
		}
	})

}
