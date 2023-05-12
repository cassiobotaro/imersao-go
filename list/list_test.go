package main

import (
	"fmt"
	"testing"
)

func BenchmarkUseContainerList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usaContainerList(10_000)
	}
}

func BenchmarkUseSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usaSlice(10_000)
	}
}

func TestComparePerformance(t *testing.T) {
	containerListTime := testing.Benchmark(BenchmarkUseContainerList).String()
	sliceTime := testing.Benchmark(BenchmarkUseSlice).String()

	fmt.Println("Container List Time: ", containerListTime)
	fmt.Println("Slice Time: ", sliceTime)
}
