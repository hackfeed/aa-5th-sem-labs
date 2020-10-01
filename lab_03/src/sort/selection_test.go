package sort

import (
	"math/rand"
	"testing"
)

// Selection sort benchmarks.

func BenchmarkSelectionSorted10(b *testing.B) {
	N := 10
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Selection(narr)
	}
}

func BenchmarkSelectionSorted100(b *testing.B) {
	N := 100
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Selection(narr)
	}
}

func BenchmarkSelectionSorted1000(b *testing.B) {
	N := 1000
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Selection(narr)
	}
}

func BenchmarkSelectionReverseSorted10(b *testing.B) {
	N := 10
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = N - i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Selection(narr)
	}
}

func BenchmarkSelectionReverseSorted100(b *testing.B) {
	N := 100
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = N - i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Selection(narr)
	}
}

func BenchmarkSelectionReverseSorted1000(b *testing.B) {
	N := 1000
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = N - i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Selection(narr)
	}
}

func BenchmarkSelectionRandom10(b *testing.B) {
	N := 10
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N-i) + i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Selection(narr)
	}
}

func BenchmarkSelectionRandom100(b *testing.B) {
	N := 100
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N-i) + i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Selection(narr)
	}
}

func BenchmarkSelectionRandom1000(b *testing.B) {
	N := 1000
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N-i) + i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Selection(narr)
	}
}
