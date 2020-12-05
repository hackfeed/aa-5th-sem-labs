package dict

import (
	"sort"
	"testing"
)

var darr = CreateArray(10000)
var farr = darr.FAnalysis()

// Brute benchmarks.
func BenchmarkBruteA(b *testing.B) {
	w := darr.Pick("a")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteB(b *testing.B) {
	w := darr.Pick("b")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteC(b *testing.B) {
	w := darr.Pick("c")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteD(b *testing.B) {
	w := darr.Pick("d")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteE(b *testing.B) {
	w := darr.Pick("e")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteF(b *testing.B) {
	w := darr.Pick("f")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteG(b *testing.B) {
	w := darr.Pick("g")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteH(b *testing.B) {
	w := darr.Pick("h")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteI(b *testing.B) {
	w := darr.Pick("i")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteJ(b *testing.B) {
	w := darr.Pick("j")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteK(b *testing.B) {
	w := darr.Pick("k")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteL(b *testing.B) {
	w := darr.Pick("l")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteM(b *testing.B) {
	w := darr.Pick("m")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteN(b *testing.B) {
	w := darr.Pick("n")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteO(b *testing.B) {
	w := darr.Pick("o")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteP(b *testing.B) {
	w := darr.Pick("p")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteQ(b *testing.B) {
	w := darr.Pick("q")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteR(b *testing.B) {
	w := darr.Pick("r")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteS(b *testing.B) {
	w := darr.Pick("s")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteT(b *testing.B) {
	w := darr.Pick("t")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteU(b *testing.B) {
	w := darr.Pick("u")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteV(b *testing.B) {
	w := darr.Pick("v")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteW(b *testing.B) {
	w := darr.Pick("w")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteX(b *testing.B) {
	w := darr.Pick("x")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteY(b *testing.B) {
	w := darr.Pick("y")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

func BenchmarkBruteZ(b *testing.B) {
	w := darr.Pick("z")
	for i := 0; i < b.N; i++ {
		darr.Brute(w)
	}
}

// Binary benchmarks
func BenchmarkBinaryA(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("a")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryB(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("b")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryC(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("c")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryD(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("d")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryE(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("e")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryF(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("f")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryG(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("g")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryH(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("h")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryI(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("i")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryJ(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("j")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryK(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("k")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryL(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("l")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryM(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("m")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryN(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("n")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryO(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("o")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryP(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("p")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryQ(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("q")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryR(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("r")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryS(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("s")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryT(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("t")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryU(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("u")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryV(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("v")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryW(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("w")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryX(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("x")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryY(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("y")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

func BenchmarkBinaryZ(b *testing.B) {
	darrc := make(DictArray, len(darr))
	copy(darrc, darr)
	sort.Slice(darrc, func(i, j int) bool {
		return darrc[i]["gamertag"].(string) < darrc[j]["gamertag"].(string)
	})
	w := darrc.Pick("z")
	for i := 0; i < b.N; i++ {
		darrc.Binary(w)
	}
}

// Combined benchmarks

func BenchmarkCombinedA(b *testing.B) {
	w := darr.Pick("a")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedB(b *testing.B) {
	w := darr.Pick("b")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedC(b *testing.B) {
	w := darr.Pick("c")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedD(b *testing.B) {
	w := darr.Pick("d")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedE(b *testing.B) {
	w := darr.Pick("e")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedF(b *testing.B) {
	w := darr.Pick("f")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedG(b *testing.B) {
	w := darr.Pick("g")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedH(b *testing.B) {
	w := darr.Pick("h")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedI(b *testing.B) {
	w := darr.Pick("i")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedJ(b *testing.B) {
	w := darr.Pick("j")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedK(b *testing.B) {
	w := darr.Pick("k")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedL(b *testing.B) {
	w := darr.Pick("l")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedM(b *testing.B) {
	w := darr.Pick("m")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedN(b *testing.B) {
	w := darr.Pick("n")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedO(b *testing.B) {
	w := darr.Pick("o")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedP(b *testing.B) {
	w := darr.Pick("p")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedQ(b *testing.B) {
	w := darr.Pick("q")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedR(b *testing.B) {
	w := darr.Pick("r")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedS(b *testing.B) {
	w := darr.Pick("s")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedT(b *testing.B) {
	w := darr.Pick("t")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedU(b *testing.B) {
	w := darr.Pick("u")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedV(b *testing.B) {
	w := darr.Pick("v")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedW(b *testing.B) {
	w := darr.Pick("w")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedX(b *testing.B) {
	w := darr.Pick("x")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedY(b *testing.B) {
	w := darr.Pick("y")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}

func BenchmarkCombinedZ(b *testing.B) {
	w := darr.Pick("z")
	for i := 0; i < b.N; i++ {
		farr.Combined(w)
	}
}
