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
	amat := matrix.ReadMatrix(an, am)

	fmt.Println(aurora.Cyan("\nМатрица B"))
	fmt.Print(aurora.Cyan("Введите количество строк   : "))
	bn := matrix.ReadNum()

	if bn != am {
		fmt.Println(aurora.Red("Количество строк матрицы B не равно количеству" +
			" столбцов матрицы A. Умножение невозможно."))
		return
	}

	fmt.Print(aurora.Cyan("Введите количество столбцов: "))
	bm := matrix.ReadNum()
	fmt.Println(aurora.Cyan("Введите матрицу:"))
	bmat := matrix.ReadMatrix(bn, bm)

	fmt.Println(aurora.Yellow("\nРезультирующая матрица"))
	fmt.Println(aurora.Cyan("Обычное умножение:"))
	smmat := matrix.SimpleMult(amat, bmat)
	smmat.PrintMatrix()
}
