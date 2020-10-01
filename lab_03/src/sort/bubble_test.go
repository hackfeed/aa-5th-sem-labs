package sort

import (
	"math/rand"
	"testing"
)

// Bubble sort benchmarks.

func BenchmarkBubbleSorted10(b *testing.B) {
	N := 10
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

func BenchmarkBubbleSorted100(b *testing.B) {
	N := 100
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

func BenchmarkBubbleSorted1000(b *testing.B) {
	N := 1000
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

func BenchmarkBubbleReverseSorted10(b *testing.B) {
	N := 10
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

func BenchmarkBubbleReverseSorted100(b *testing.B) {
	N := 100
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

func BenchmarkBubbleReverseSorted1000(b *testing.B) {
	N := 1000
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

func BenchmarkBubbleRandom10(b *testing.B) {
	N := 10
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

func BenchmarkBubbleRandom100(b *testing.B) {
	N := 100
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

func BenchmarkBubbleRandom1000(b *testing.B) {
	N := 1000
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
