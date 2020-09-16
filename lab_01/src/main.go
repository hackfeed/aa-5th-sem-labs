package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"

	"./levenshtein"
)

func main() {
	fmt.Println(aurora.Magenta("Расстояние Левенштейна\n"))

	fmt.Printf("Введите первое слово: ")
	fWord := levenshtein.ReadWord()
	fmt.Printf("Введите второе слово: ")
	sWord := levenshtein.ReadWord()

	fmt.Println()

	distRec := levenshtein.Recursive(fWord, sWord)
	fmt.Println(aurora.Yellow("Рекурсивный метод без заполнения матрицы:"))
	fmt.Printf("Расстояние: %d\n\n", aurora.Green(distRec))

	fmt.Println()

	distRM, matRM := levenshtein.RecursiveMatrix(fWord, sWord)
	fmt.Println(aurora.Yellow("Рекурсивный метод с заполнением матрицы:"))
	matRM.PrintMatrix()
	fmt.Printf("Расстояние: %d\n\n", aurora.Green(distRM))

	fmt.Println()

	distIt, matIt := levenshtein.IterativeMatrix(fWord, sWord)
	fmt.Println(aurora.Yellow("Итеративный метод с заполнением матрицы:"))
	matIt.PrintMatrix()
	fmt.Printf("Расстояние: %d\n\n", aurora.Green(distIt))

	fmt.Println()

	distDL, matDL := levenshtein.DamerauLevenshtein(fWord, sWord)
	fmt.Println(aurora.Yellow("Метод Дамерау - Левенштейна:"))
	matDL.PrintMatrix()
	fmt.Printf("Расстояние: %d\n\n", aurora.Green(distDL))
}
