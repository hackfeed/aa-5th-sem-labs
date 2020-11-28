package conv

import (
	"time"

	"github.com/brianvoe/gofakeit"
)

// Conveyor used to generate credit cards in parallel conveyor way.
func Conveyor(n int, ch chan int) *Queue {
	f := make(chan *CreditCard, 5)
	s := make(chan *CreditCard, 5)
	t := make(chan *CreditCard, 5)
	l := createQueue(n)

	number := func() {
		for {
			select {
			case cc := <-f:
				cc.haveNumber = true

				cc.startNumber = time.Now()
				cc.number = gofakeit.CreditCardNumber(nil)
				cc.doneNumber = time.Now()

				s <- cc
			}
		}
	}

	cvv := func() {
		for {
			select {
			case cc := <-s:
				cc.haveCvv = true

				cc.startCvv = time.Now()
				cc.cvv = gofakeit.CreditCardCvv()
				cc.doneCvv = time.Now()

				t <- cc
			}
		}
	}

	exp := func() {
		for {
			select {
			case cc := <-t:
				cc.haveExp = true

				cc.startExp = time.Now()
				cc.exp = gofakeit.CreditCardExp()
				cc.doneExp = time.Now()

				l.push(cc)
				if cc.num == n {
					ch <- 0
				}

			}
		}
	}

	go number()
	go cvv()
	go exp()
	for i := 0; i <= n; i++ {
		cc := new(CreditCard)
		cc.num = i
		f <- cc
	}

	return l
}

func number(cc *CreditCard, qNumber *Queue) {
	cc.haveNumber = true

	cc.startNumber = time.Now()
	cc.number = gofakeit.CreditCardNumber(nil)
	cc.doneNumber = time.Now()

	qNumber.push(cc)
}

func cvv(cc *CreditCard, qCvv *Queue) {
	cc.haveCvv = true

	cc.startCvv = time.Now()
	cc.cvv = gofakeit.CreditCardCvv()
	cc.doneCvv = time.Now()

	qCvv.push(cc)
}

func exp(cc *CreditCard, done *Queue) {
	cc.haveExp = true

	cc.startExp = time.Now()
	cc.exp = gofakeit.CreditCardExp()
	cc.doneExp = time.Now()

	done.push(cc)
}
