package sort

import (
	"math/rand"
	"testing"
)

var N = 1000

// Bubble sort benchmarks.

func BenchmarkBubbleSorted(b *testing.B) {
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Bubble(narr)
	}
}

func BenchmarkBubbleReverseSorted(b *testing.B) {
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = N - i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Bubble(narr)
	}
}

func BenchmarkBubbleRandom(b *testing.B) {
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N-i) + i
	}

	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Bubble(narr)
	}
}

// Insertion sort benchmarks.

func BenchmarkInsertionSorted(b *testing.B) {
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

func BenchmarkInsertionReverseSorted(b *testing.B) {
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

func BenchmarkInsertionRandom(b *testing.B) {
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

// Selection sort becnhmarks.

func BenchmarkSelectionSorted(b *testing.B) {
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

func BenchmarkSelectionReverseSorted(b *testing.B) {
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

func BenchmarkSelectionRandom(b *testing.B) {
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
