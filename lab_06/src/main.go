package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"

	"./ant"
)

func main() {
	fmt.Printf("%v", aurora.Magenta("Муравьиный алгоритм\n\n"))

	ant.Compare("data/data.txt")
}
