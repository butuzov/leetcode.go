package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func tests(t *testing.T, fnc func(int) int) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("MySqrt(%d)", i), func(t *testing.T) {
			assert.Equal(t, test.expected, fnc(test.x), "Input(x=(%d))", test.x)
		})
	}
}

func becnhmarks(b *testing.B, fnc func(int) int) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				fnc(test.x)
			}
		}
	})
}

func TestBabylonian(t *testing.T) {
	tests(t, Babylonian)
}

func TestBakhshali(t *testing.T) {
	tests(t, Bakhshali)
}

func TestBits(t *testing.T) {
	tests(t, Bits)
}

func BenchmarkBabylonian(b *testing.B) {
	becnhmarks(b, Babylonian)
}

func BenchmarkBakhshali(b *testing.B) {
	becnhmarks(b, Bakhshali)
}
func BenchmarkBits(b *testing.B) {
	becnhmarks(b, Bits)
}
