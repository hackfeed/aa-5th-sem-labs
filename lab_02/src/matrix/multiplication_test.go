package matrix

import (
	"testing"
)

// Simple multiplication benchmarks.

func BenchmarkSimple100(b *testing.B) {
	N := 100
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

func BenchmarkSimple200(b *testing.B) {
	N := 200
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

func BenchmarkSimple300(b *testing.B) {
	N := 300
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

func BenchmarkSimple400(b *testing.B) {
	N := 400
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

func BenchmarkSimple500(b *testing.B) {
	N := 500
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

// Winograd multiplication benchmarks.

func BenchmarkWinograd100(b *testing.B) {
	N := 100
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

func BenchmarkWinograd200(b *testing.B) {
	N := 200
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

func BenchmarkWinograd300(b *testing.B) {
	N := 300
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

func BenchmarkWinograd400(b *testing.B) {
	N := 400
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

func BenchmarkWinograd500(b *testing.B) {
	N := 500
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

// Winograd improved multiplication benchmarks.

func BenchmarkWinogradImp100(b *testing.B) {
	N := 100
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}

func BenchmarkWinogradImp200(b *testing.B) {
	N := 200
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}

func BenchmarkWinogradImp300(b *testing.B) {
	N := 300
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}

func BenchmarkWinogradImp400(b *testing.B) {
	N := 400
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}

func BenchmarkWinogradImp500(b *testing.B) {
	N := 500
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}
