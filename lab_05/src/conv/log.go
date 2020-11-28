package conv

import (
	"fmt"
	"strings"
	"time"

	"github.com/logrusorgru/aurora"
)

// Log used to log conveyor tasks time.
func Log(q *Queue, logCreditCard bool) {
	var (
		fdtime time.Duration
		sdtime time.Duration
		tdtime time.Duration
	)

	l := q.q
	start := l[0].startNumber

	fmt.Printf("%38v\n", aurora.BgRed("STARTING TIME"))
	fmt.Printf("%v%36v%v\n", "+", strings.Repeat("-", 36), "+")
	fmt.Printf(
		"|%3v|%10v|%10v|%10v|\n",
		aurora.Green("N"), aurora.Green("CardNumber"), aurora.Green("CardCVV"), aurora.Green("CardExp"),
	)
	fmt.Printf("%v%36v%v\n", "+", strings.Repeat("-", 36), "+")
	for i := 0; i < len(l); i++ {
		if l[i] != nil {
			fmt.Printf(
				"|%3v|%10v|%10v|%10v|\n",
				i, l[i].startNumber.Sub(start), l[i].startCvv.Sub(start), l[i].startExp.Sub(start),
			)
		}
	}
	fmt.Printf("%v%36v%v\n", "+", strings.Repeat("-", 36), "+")

	fmt.Printf("%38v\n", aurora.BgRed("FINISHING TIME"))
	fmt.Printf("%v%36v%v\n", "+", strings.Repeat("-", 36), "+")
	fmt.Printf(
		"|%3v|%10v|%10v|%10v|\n",
		aurora.Green("N"), aurora.Green("CardNumber"), aurora.Green("CardCVV"), aurora.Green("CardExp"),
	)
	fmt.Printf("%v%36v%v\n", "+", strings.Repeat("-", 36), "+")
	for i := 0; i < len(l); i++ {
		if l[i] != nil {
			fmt.Printf(
				"|%3v|%10v|%10v|%10v|\n",
				i, l[i].doneNumber.Sub(start), l[i].doneCvv.Sub(start), l[i].doneExp.Sub(start),
			)
		}
	}
	fmt.Printf("%v%36v%v\n", "+", strings.Repeat("-", 36), "+")

	fmt.Printf("%38v\n", aurora.BgRed("IDLE TIME"))
	fmt.Printf("%v%36v%v\n", "+", strings.Repeat("-", 36), "+")
	fmt.Printf("|%3v|%32v|\n", aurora.Green("N"), aurora.Green("Idle time"))
	fmt.Printf("%v%36v%v\n", "+", strings.Repeat("-", 36), "+")
	for i := 0; i < len(l)-1; i++ {
		fdtime += l[i+1].startNumber.Sub(start) - l[i].doneNumber.Sub(start)
		sdtime += l[i+1].startCvv.Sub(start) - l[i].doneCvv.Sub(start)
		tdtime += l[i+1].startExp.Sub(start) - l[i].doneExp.Sub(start)
	}
	ldtime := []time.Duration{fdtime, sdtime, tdtime}
	for i := 0; i < 3; i++ {
		fmt.Printf("|%3v|%32v|\n", i+1, ldtime[i])
	}
	fmt.Printf("%v%36v%v\n", "+", strings.Repeat("-", 36), "+")

	if logCreditCard {
		fmt.Printf("\n%v:\n", aurora.Magenta("Generated data"))
		for i := range l {
			fmt.Printf("CreditCard: %d\nNumber: %s\nCVV: %s\nExpiration date: %s\n\n", l[i].num+1, l[i].number, l[i].cvv, l[i].exp)
		}
	}
}
