package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"

	"./matrix"
)

func main() {
	fmt.Println(aurora.Magenta("Умножение матриц\n"))

	fmt.Println(aurora.Cyan("Матрица А"))
	fmt.Print(aurora.Cyan("Введите количество строк   : "))
	an := matrix.ReadNum()
	fmt.Print(aurora.Cyan("Введите количество столбцов: "))
	am := matrix.ReadNum()
	fmt.Println(aurora.Cyan("Введите матрицу:"))
	_ = matrix.ReadMatrix(an, am)

	fmt.Println(aurora.Cyan("\nМатрица B"))
	fmt.Print(aurora.Cyan("Введите количество строк   : "))
	bn := matrix.ReadNum()
	fmt.Print(aurora.Cyan("Введите количество столбцов: "))
	bm := matrix.ReadNum()
	fmt.Println(aurora.Cyan("Введите матрицу:"))
	_ = matrix.ReadMatrix(bn, bm)
}
