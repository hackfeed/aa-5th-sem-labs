package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"

	"./levenshtein"
)

func main() {
	fmt.Println(aurora.Magenta("Levenshtein distance finder\n"))

	fmt.Printf("Input first word : ")
	fWord := levenshtein.ReadWord()
	fmt.Printf("Input second word: ")
	sWord := levenshtein.ReadWord()

	fmt.Println()

	distRec := levenshtein.Recursive(fWord, sWord)
	fmt.Println(aurora.Yellow("Recursive method without matrix filling: "))
	fmt.Printf("Distance: %d\n\n", aurora.Green(distRec))

	fmt.Println()

	distRM, matRM := levenshtein.RecursiveMatrix(fWord, sWord)
	fmt.Println(aurora.Yellow("Recursive method with matrix filling: "))
	matRM.PrintMatrix()
	fmt.Printf("Distance: %d\n\n", aurora.Green(distRM))

	fmt.Println()

	distIt, matIt := levenshtein.IterativeMatrix(fWord, sWord)
	fmt.Println(aurora.Yellow("Iterative method with matrix filling: \n"))
	matIt.PrintMatrix()
	fmt.Printf("Distance: %d\n\n", aurora.Green(distIt))

	fmt.Println()

	distDL, matDL := levenshtein.DamerauLevenshtein(fWord, sWord)
	fmt.Println(aurora.Yellow("Damerau-Levenshtein method with matrix filling: \n"))
	matDL.PrintMatrix()
	fmt.Printf("Distance: %d\n\n", aurora.Green(distDL))
}
