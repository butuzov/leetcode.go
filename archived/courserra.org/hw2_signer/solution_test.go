package main

import "testing"

var benchResult string

func bench(b *testing.B, fnc func(string) string) {

	var inputData = []string{"0", "1", "1", "2", "3", "5", "8"}

	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r string
			for i := range inputData {
				r = fnc(inputData[i])
			}
			benchResult = r
		}
	})
}

func BenchmarkNative(b *testing.B) {
	bench(b, DataSignerMd5)
}

func BenchmarkRatelimiter(b *testing.B) {
	bench(b, Signer(DataSignerMd5))
}
