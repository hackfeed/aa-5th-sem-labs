package levenshtein

import (
	"testing"
)

// Recursive method benchmarks.

func BenchmarkRecursiveLen5(b *testing.B) {
	fWord := "about"
	sWord := "above"

	for i := 0; i < b.N; i++ {
		Recursive(fWord, sWord)
	}
}

func BenchmarkRecursiveLen10(b *testing.B) {
	fWord := "abbanition"
	sWord := "abaptiston"

	for i := 0; i < b.N; i++ {
		Recursive(fWord, sWord)
	}
}

func BenchmarkRecursiveLen15(b *testing.B) {
	fWord := "characteristics"
	sWord := "recommendations"

	for i := 0; i < b.N; i++ {
		Recursive(fWord, sWord)
	}
}

// RecursiveMatrix method becnhmarks.

func BenchmarkRecursiveMatrixLen10(b *testing.B) {
	fWord := "abbanition"
	sWord := "abaptiston"

	for i := 0; i < b.N; i++ {
		RecursiveMatrix(fWord, sWord)
	}
}

func BenchmarkRecursiveMatrixLen20(b *testing.B) {
	fWord := "abdominohysterectomy"
	sWord := "acetylcholinesterase"

	for i := 0; i < b.N; i++ {
		RecursiveMatrix(fWord, sWord)
	}
}
func BenchmarkRecursiveMatrixLen30(b *testing.B) {
	fWord := "chlorobenzylidenemalononitrile"
	sWord := "abdominalexternalobliquemuscle"

	for i := 0; i < b.N; i++ {
		RecursiveMatrix(fWord, sWord)
	}
}

func BenchmarkRecursiveMatrixLen50(b *testing.B) {
	fWord := "chlorobenzylidenemalononitrileabdominohysterectomy"
	sWord := "abdominalexternalobliquemuscleacetylcholinesterase"

	for i := 0; i < b.N; i++ {
		RecursiveMatrix(fWord, sWord)
	}
}

func BenchmarkRecursiveMatrixLen100(b *testing.B) {
	fWord := "chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy"
	sWord := "abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase"

	for i := 0; i < b.N; i++ {
		RecursiveMatrix(fWord, sWord)
	}
}

func BenchmarkRecursiveMatrixLen200(b *testing.B) {
	fWord := "chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy"
	sWord := "abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase"

	for i := 0; i < b.N; i++ {
		RecursiveMatrix(fWord, sWord)
	}
}

// IterativeMatrix method benchmarks.

func BenchmarkIterativeMatrixLen10(b *testing.B) {
	fWord := "abbanition"
	sWord := "abaptiston"

	for i := 0; i < b.N; i++ {
		IterativeMatrix(fWord, sWord)
	}
}

func BenchmarkIterativeMatrixLen20(b *testing.B) {
	fWord := "abdominohysterectomy"
	sWord := "acetylcholinesterase"

	for i := 0; i < b.N; i++ {
		IterativeMatrix(fWord, sWord)
	}
}

func BenchmarkIterativeMatrixLen30(b *testing.B) {
	fWord := "chlorobenzylidenemalononitrile"
	sWord := "abdominalexternalobliquemuscle"

	for i := 0; i < b.N; i++ {
		IterativeMatrix(fWord, sWord)
	}
}

func BenchmarkIterativeMatrixLen50(b *testing.B) {
	fWord := "chlorobenzylidenemalononitrileabdominohysterectomy"
	sWord := "abdominalexternalobliquemuscleacetylcholinesterase"

	for i := 0; i < b.N; i++ {
		IterativeMatrix(fWord, sWord)
	}
}

func BenchmarkIterativeMatrixLen100(b *testing.B) {
	fWord := "chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy"
	sWord := "abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase"

	for i := 0; i < b.N; i++ {
		IterativeMatrix(fWord, sWord)
	}
}

func BenchmarkIterativeMatrixLen200(b *testing.B) {
	fWord := "chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy"
	sWord := "abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase"

	for i := 0; i < b.N; i++ {
		IterativeMatrix(fWord, sWord)
	}
}

// DamerauLevenshtein method benchmarks.

func BenchmarkDamerauLevenshteinLen10(b *testing.B) {
	fWord := "abbanition"
	sWord := "abaptiston"

	for i := 0; i < b.N; i++ {
		DamerauLevenshtein(fWord, sWord)
	}
}

func BenchmarkDamerauLevenshteinLen20(b *testing.B) {
	fWord := "abdominohysterectomy"
	sWord := "acetylcholinesterase"

	for i := 0; i < b.N; i++ {
		DamerauLevenshtein(fWord, sWord)
	}
}

func BenchmarkDamerauLevenshteinLen30(b *testing.B) {
	fWord := "chlorobenzylidenemalononitrile"
	sWord := "abdominalexternalobliquemuscle"

	for i := 0; i < b.N; i++ {
		DamerauLevenshtein(fWord, sWord)
	}
}

func BenchmarkDamerauLevenshteinLen50(b *testing.B) {
	fWord := "chlorobenzylidenemalononitrileabdominohysterectomy"
	sWord := "abdominalexternalobliquemuscleacetylcholinesterase"

	for i := 0; i < b.N; i++ {
		DamerauLevenshtein(fWord, sWord)
	}
}

func BenchmarkDamerauLevenshteinLen100(b *testing.B) {
	fWord := "chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy"
	sWord := "abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase"

	for i := 0; i < b.N; i++ {
		DamerauLevenshtein(fWord, sWord)
	}
}

func BenchmarkDamerauLevenshteinLen200(b *testing.B) {
	fWord := "chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy"
	sWord := "abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase"

	for i := 0; i < b.N; i++ {
		DamerauLevenshtein(fWord, sWord)
	}
}
