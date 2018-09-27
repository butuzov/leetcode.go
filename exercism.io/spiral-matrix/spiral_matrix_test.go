package spiralmatrix

import (
	"reflect"
	"testing"
)

func TestSpiralMatrix(t *testing.T) {
	for _, testCase := range testCases {
		matrix := SpiralMatrix(testCase.input)
		if !reflect.DeepEqual(matrix, testCase.expected) {
			t.Fatalf("FAIL: %s\n\tSpiralMatrix(%v)\nexpected: %v\ngot     : %v",
				testCase.description, testCase.input, testCase.expected, matrix)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

func BenchmarkSpiralMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			SpiralMatrix(testCase.input)
		}
	}
}
