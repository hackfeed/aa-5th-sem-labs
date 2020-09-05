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
	_, mat := levenshtein.IterativeMatrix(fWord, sWord)

	mat.PrintMatrix()
}
