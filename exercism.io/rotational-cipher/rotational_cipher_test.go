package rotationalcipher

import (
	"testing"
)

func TestRotationalCipher(t *testing.T) {
	for _, testCase := range testCases {
		cipher := RotationalCipher(testCase.inputPlain, testCase.inputShiftKey)
		if cipher != testCase.expected {
			t.Fatalf("FAIL: %s\n\tRotationalCipher(%s, %d)\nexpected: %s, \ngot:      %s",
				testCase.description, testCase.inputPlain, testCase.inputShiftKey, testCase.expected, cipher)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

func BenchmarkRotationalCipher(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			RotationalCipher(testCase.inputPlain, testCase.inputShiftKey)
		}
	}
}
