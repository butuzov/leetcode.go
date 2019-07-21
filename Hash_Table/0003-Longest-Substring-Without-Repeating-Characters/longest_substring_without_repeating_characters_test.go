package main

import (
	"fmt"
	"testing"
)

var MsgError = "Fail: Input(s=(%+v)): Expected(%+v) vs Returend(%+v)"
var MsgOk = "OK: Input(s=(%+v)): as Expected(%+v)"

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		test := test

		t.Run(fmt.Sprintf("LengthOfLongestSubstring(%d)", i), func(t *testing.T) {
			res := lengthOfLongestSubstring(test.s)

			// reflect.DeepEqual(res, test.expected)
			if res != test.expected {
				t.Errorf(MsgError, test.s, test.expected, res)
				return
			}
			t.Logf(MsgOk, test.s, test.expected)
		})
	}
}

var result int

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var r int
		for _, test := range TestCases {
			r = lengthOfLongestSubstring(test.s)
		}
		result = r
	}
}
