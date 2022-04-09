package benchmarks

import (
	"math/rand"
	"sort"
	"testing"
)

func genIntSlice(n int) []int {
	genFloatSlice(10)
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, rand.Intn(1e9))
	}
	return s
}

func genFloatSlice(n int) []float64 {
	s := make([]float64, 0, n)

	for i := 0; i < n; i++ {
		s = append(s, rand.Float64())
	}

	return s
}

func BenchmarkSortInt10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := genIntSlice(10000)
		b.StartTimer()
		sort.Ints(s)
	}
}

func BenchmarkSortFloat10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := genFloatSlice(10000)
		b.StartTimer()
		sort.Float64s(s)
	}
}

func BenchmarkSortInt100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := genIntSlice(100000)
		b.StartTimer()
		sort.Ints(s)
	}
}

func BenchmarkSortFloat100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := genFloatSlice(100000)
		b.StartTimer()
		sort.Float64s(s)
	}
}

func BenchmarkSortInt1000000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := genIntSlice(1000000)
		b.StartTimer()
		sort.Ints(s)
	}
}

func BenchmarkSortFloat1000000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := genFloatSlice(1000000)
		b.StartTimer()
		sort.Float64s(s)
	}
}

// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
//
// BenchmarkSortInt10000-12                    1262            900734 ns/op
// BenchmarkSortFloat10000-12                  1160           1025616 ns/op
// BenchmarkSortInt100000-12                    100          11064761 ns/op
// BenchmarkSortFloat100000-12                   96          12576909 ns/op
// BenchmarkSortInt1000000-12                     8         130038541 ns/op
// BenchmarkSortFloat1000000-12                   7         147019343 ns/op
// PASS
// ok      command-line-arguments  8.906s