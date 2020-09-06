package main

import (
	"fmt"

	"./levenshtein"
)

func main() {
	fmt.Printf("Input first word : ")
	fWord := levenshtein.ReadWord()
	fmt.Printf("Input second word: ")
	sWord := levenshtein.ReadWord()

	fmt.Println()

	distIt, matIt := levenshtein.IterativeMatrix(fWord, sWord)
	fmt.Printf("Iterative method with matrix filling: \n")
	matIt.PrintMatrix()
	fmt.Printf("Distance: %d\n\n", distIt)

	fmt.Println()

	distDL, matDL := levenshtein.DamerauLevenshtein(fWord, sWord)
	fmt.Printf("Damerau-Levenshtein method with matrix filling: \n")
	matDL.PrintMatrix()
	fmt.Printf("Distance: %d\n\n", distDL)

	fmt.Println()

	distRec := levenshtein.Recursive(fWord, sWord)
	fmt.Printf("Recursive method without matrix filling: \n")
	fmt.Printf("Distance: %d\n\n", distRec)
}
