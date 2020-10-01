package sort

import (
	"math/rand"
	"testing"
)

// Insertion sort benchmarks.

func BenchmarkInsertionSorted10(b *testing.B) {
	N := 10
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Insertion(narr)
	}
}

func BenchmarkInsertionSorted100(b *testing.B) {
	N := 100
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Insertion(narr)
	}
}

func BenchmarkInsertionSorted1000(b *testing.B) {
	N := 1000
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Insertion(narr)
	}
}

func BenchmarkInsertionReverseSorted10(b *testing.B) {
	N := 10
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = N - i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Insertion(narr)
	}
}

func BenchmarkInsertionReverseSorted100(b *testing.B) {
	N := 100
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = N - i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Insertion(narr)
	}
}

func BenchmarkInsertionReverseSorted1000(b *testing.B) {
	N := 1000
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = N - i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Insertion(narr)
	}
}

func BenchmarkInsertionRandom10(b *testing.B) {
	N := 10
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N-i) + i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Insertion(narr)
	}
}

func BenchmarkInsertionRandom100(b *testing.B) {
	N := 100
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N-i) + i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Insertion(narr)
	}
}

func BenchmarkInsertionRandom1000(b *testing.B) {
	N := 1000
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N-i) + i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Insertion(narr)
	}
}
