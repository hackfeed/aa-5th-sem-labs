package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"

	labsort "./sort"
)

func main() {
	fmt.Println(aurora.Magenta("Методы сортировки массивов"))

	fmt.Print(aurora.Cyan("Введите длину массива: "))
	n := labsort.ReadNum()
	fmt.Print(aurora.Cyan("Введите массив: "))
	arr := labsort.ReadArray(n)
	barr := make([]int, n)
	iarr := make([]int, n)
	sarr := make([]int, n)

	fmt.Print(aurora.Yellow("Сортировка пузырьком: "))
	copy(barr, arr)
	labsort.Bubble(barr)
	fmt.Println(barr)

	fmt.Print(aurora.Yellow("Сортировка вставками: "))
	copy(iarr, arr)
	labsort.Insertion(iarr)
	fmt.Println(iarr)

	fmt.Print(aurora.Yellow("Сортировка выбором  : "))
	copy(sarr, arr)
	labsort.Selection(sarr)
	fmt.Println(sarr)
}
