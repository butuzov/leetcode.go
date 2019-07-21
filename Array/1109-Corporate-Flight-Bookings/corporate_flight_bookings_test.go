package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorpFlightBookings(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("CorpFlightBookings(%d)", i), func(t *testing.T) {
			assert.Equal(t, test.expected, corpFlightBookings(test.bookings, test.n))
		})
	}
}

func BenchmarkCorpFlightBookings(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				corpFlightBookings(test.bookings, test.n)
			}
		}
	})
}
