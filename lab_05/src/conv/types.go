package conv

import "time"

type CreditCard struct {
	num         int
	haveNumber  bool
	haveCvv     bool
	haveExp     bool
	startNumber time.Time
	doneNumber  time.Time
	startCvv    time.Time
	doneCvv     time.Time
	startExp    time.Time
	doneExp     time.Time

	number string
	cvv    string
	exp    string
}

type Queue struct {
	q [](*CreditCard)
	l int
}
