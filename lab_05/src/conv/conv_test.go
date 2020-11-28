package conv

import "testing"

func BenchmarkConveyor10(b *testing.B) {
	N := 10
	ch := make(chan int)

	for i := 0; i < b.N; i++ {
		Conveyor(N, ch)
		<-ch
	}
}

func BenchmarkConveyor20(b *testing.B) {
	N := 20

	for i := 0; i < b.N; i++ {
		ch := make(chan int)
		Conveyor(N, ch)
		<-ch
	}
}

func BenchmarkConveyor30(b *testing.B) {
	N := 30

	for i := 0; i < b.N; i++ {
		ch := make(chan int)
		Conveyor(N, ch)
		<-ch
	}
}

func BenchmarkConveyor40(b *testing.B) {
	N := 40

	for i := 0; i < b.N; i++ {
		ch := make(chan int)
		Conveyor(N, ch)
		<-ch
	}
}

func BenchmarkConveyor50(b *testing.B) {
	N := 50

	for i := 0; i < b.N; i++ {
		ch := make(chan int)
		Conveyor(N, ch)
		<-ch
	}
}
