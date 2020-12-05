package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"

	"./dict"
)

func main() {
	fmt.Printf("%v", aurora.Magenta("Поиск в массиве словарей\n\n"))

	darr := dict.CreateArray(10)
	farr := darr.FAnalysis()
	gt := "applicationpursue849"

	fmt.Printf("%v %v\n\n", aurora.Green("Ключ для поиска:"), aurora.Blue(gt))
	darr.Print()
	fmt.Println()

	r := farr.Combined(gt)
	if r["gamertag"] == nil {
		fmt.Printf("%v\n", aurora.Red("Словарь с данным ключом не найден"))
	} else {
		fmt.Printf("%v\n", aurora.Green("Словарь с данным ключом найден"))
		r.Print()
	}
}
