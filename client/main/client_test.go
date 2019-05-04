package main

import (
	"math/rand"
	"testing"
)

func BenchmarkDoJobClient(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		xdata := data{Key: i, Value: rand.Intn(100000)}
		job := Job{data: xdata}
		job.data.DoJob()
	}
}
