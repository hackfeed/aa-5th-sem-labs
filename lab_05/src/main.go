package main

import (
	"fmt"
	"time"

	"github.com/logrusorgru/aurora"

	"./conv"
)

func main() {
	fmt.Printf("%v", aurora.Magenta("Параллельные конвейерные вычисления\n\n"))

	n := 20

	start := time.Now()
	ch := make(chan int)
	lineP := conv.Conveyor(n, ch)
	<-ch
	conv.Log(lineP, false)
	_ = lineP
	end := time.Now()
	fmt.Println(
		aurora.Sprintf(
			aurora.BrightGreen("\nОформление %d кредитных кард на конвейере заняло %v\n"),
			aurora.Magenta(n),
			aurora.Magenta(end.Sub(start)),
		))
}
